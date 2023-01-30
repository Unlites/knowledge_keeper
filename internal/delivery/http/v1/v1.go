package v1

import (
	"github.com/Unlites/knowledge_keeper/internal/usecases"
	"github.com/Unlites/knowledge_keeper/pkg/logger"
	"github.com/gin-gonic/gin"
)

type v1Handler struct {
	usecases *usecases.Usecases
	log      logger.Logger
	v1Group  *gin.RouterGroup
}

func NewV1Handler(usecases *usecases.Usecases, log logger.Logger, v1Group *gin.RouterGroup) *v1Handler {
	return &v1Handler{
		usecases: usecases,
		log:      log,
		v1Group:  v1Group,
	}
}

func (h *v1Handler) Init() {
	h.v1Group.Use(h.addLoggerRequestId)

	authGroup := h.v1Group.Group("/auth")
	recordGroup := h.v1Group.Group("/record", h.userIdentification, h.addLoggerUserId)

	h.initAuthRoutes(authGroup)
	h.initRecordRoutes(recordGroup)
}
