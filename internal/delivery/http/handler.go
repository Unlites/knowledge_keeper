package http

import (
	v1 "github.com/Unlites/knowledge_keeper/internal/delivery/http/v1"
	"github.com/Unlites/knowledge_keeper/internal/usecases"
	"github.com/Unlites/knowledge_keeper/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	usecases *usecases.Usecases
	log      logger.Logger
	router   *gin.Engine
}

func NewHandler(usecases *usecases.Usecases, log logger.Logger, router *gin.Engine) *Handler {
	return &Handler{
		usecases: usecases,
		log:      log,
		router:   router,
	}
}

func (h *Handler) InitAPI() {
	api := h.router.Group("/api")
	v1Group := api.Group("/v1")

	v1.NewV1Handler(h.usecases, h.log, v1Group).Init()
}
