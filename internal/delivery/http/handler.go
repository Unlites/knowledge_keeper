package delivery

import (
	v1 "github.com/Unlites/knowledge_keeper/internal/delivery/http/v1"
	"github.com/Unlites/knowledge_keeper/internal/usecases"
	"github.com/gin-gonic/gin"
)

type handlingRoutes interface {
	InitRoutes()
}

type Handler struct {
	usecases *usecases.Usecases
	router   *gin.Engine
}

func NewHandler(usecases *usecases.Usecases, router *gin.Engine) *Handler {
	return &Handler{
		usecases: usecases,
		router:   router,
	}
}

func (h *Handler) InitAPI(router *gin.Engine) {
	api := router.Group("/api")
	v1Group := api.Group("/v1")

	authGroup := v1Group.Group("/auth")
	recordGroup := v1Group.Group("/record")

	v1.NewAuthHandler(authGroup, h.usecases.Auth).InitRoutes()
	v1.NewRecordHandler(recordGroup, h.usecases.Record).InitRoutes()
}
