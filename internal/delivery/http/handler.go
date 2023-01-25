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

func (h *Handler) InitAPI(router *gin.Engine) {
	api := router.Group("/api")
	v1Group := api.Group("/v1")

	authGroup := v1Group.Group("/auth", h.addLoggerRequestId)
	recordGroup := v1Group.Group("/record", h.addLoggerRequestId, h.addLoggerUserId)

	v1.NewAuthHandler(h.log, authGroup, h.usecases.Auth).InitRoutes()
	v1.NewRecordHandler(h.log, recordGroup, h.usecases.Record).InitRoutes()
}
