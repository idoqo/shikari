package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"gitlab.com/idoko/shikari/db"
)

type Handler struct {
	db db.Database
	logger zerolog.Logger
}

func New(db db.Database, logger zerolog.Logger) *Handler {
	return &Handler{
		db: db,
		logger: logger,
	}
}

func (h *Handler) Register(group *gin.RouterGroup) {
	group.POST("/auth/signup", h.signup)
	group.POST("/auth/login", h.login)
}