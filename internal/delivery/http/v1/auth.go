package v1

import (
	"github.com/Unlites/knowledge_keeper/internal/usecases"
	"github.com/gin-gonic/gin"
)

type authHandler struct {
	group       *gin.RouterGroup
	authUsecase usecases.Auth
}

func NewAuthHandler(group *gin.RouterGroup, authUsecase usecases.Auth) *authHandler {
	return &authHandler{
		group:       group,
		authUsecase: authUsecase,
	}
}

func (a *authHandler) InitRoutes() {

}
