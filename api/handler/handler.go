package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"gitlab.com/idoko/shikari/api/tokens"
	"gitlab.com/idoko/shikari/db"
)

type Handler struct {
	db     db.Database
	logger zerolog.Logger
	jwt    tokens.JWT
}

func New(db db.Database, logger zerolog.Logger, jwt tokens.JWT) *Handler {
	return &Handler{
		db: db,
		logger: logger,
		jwt: jwt,
	}
}

func (h *Handler) Register(group *gin.RouterGroup) {
	group.POST("/auth/signup", h.signup)
	group.POST("/auth/login", h.login)
}