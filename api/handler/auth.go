package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signup(c *gin.Context) {
	h.logger.Print("trying to signup!")
	c.JSON(http.StatusInternalServerError, gin.H{"error": "unimplemented"})
}

func (h *Handler) login(c *gin.Context) {
	h.logger.Print("trying to log in")
	c.JSON(http.StatusInternalServerError, gin.H{"error": "unimplemented"})
}