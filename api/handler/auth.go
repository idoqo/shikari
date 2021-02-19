package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.com/idoko/shikari/db"
	"gitlab.com/idoko/shikari/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func (h *Handler) signup(c *gin.Context) {
	var userReq models.UserRequest
	var err error

	if err = c.ShouldBindJSON(&userReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid request data: %s", err.Error())})
		return
	}

	var hash []byte
	hash, err = bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
	if err != nil {
		h.logger.Err(err).Msg("error while generating password hash")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to complete signup"})
		return
	}
	user := models.User{
		Email: userReq.Email,
		Password: string(hash),
	}

	_, err = h.db.GetUserByEmail(user.Email)
	if err != nil && err == db.ErrNoRecord {
		err = h.db.SaveUser(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to save user: %s", err.Error())})
			return
		}
		token, err := h.jwt.GenerateToken(user)
		if err != nil {
			h.logger.Err(err).Msg("could not generate token")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "automatic sign in failed"})
			return
		}
		res := map[string]string{
			"status": "successful",
			"token": token,
		}
		c.JSON(http.StatusOK, gin.H{"data": res})
	} else {
		if err == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("email [%s] already exists", user.Email)})
			return
		}
		h.logger.Err(err).Msg("could not complete email check")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to complete signup"})
		return
	}
}

func (h *Handler) login(c *gin.Context) {
	var userReq models.UserRequest

	if err := c.ShouldBindJSON(&userReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid request data: %s", err.Error())})
		return
	}
	user, err := h.db.GetUserByEmail(userReq.Email)
	if err != nil {
		if err == db.ErrNoRecord {
			c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("no account with provided email")})
			return
		}
		h.logger.Err(err).Msg("could not log in user")
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("could not sign in at the moment")})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userReq.Password)); err != nil {
		switch err {
		case bcrypt.ErrMismatchedHashAndPassword:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "password is incorrect"})
			break
		default:
			h.logger.Err(err).Msg("could not validate user password")
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("could not sign in at the moment")})
		}
		return
	}
	token, err := h.jwt.GenerateToken(user)
	if err != nil {
		h.logger.Err(err).Msg("could not generate token")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "automatic sign in failed"})
		return
	}
	res := map[string]string{
		"status": "successful",
		"token": token,
	}
	c.JSON(http.StatusOK, gin.H{"data": res})
}