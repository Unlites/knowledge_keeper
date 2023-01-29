package v1

import (
	"errors"
	"net/http"

	"github.com/Unlites/knowledge_keeper/internal/dto"
	"github.com/Unlites/knowledge_keeper/internal/errs"
	"github.com/Unlites/knowledge_keeper/internal/usecases"
	"github.com/Unlites/knowledge_keeper/pkg/logger"
	"github.com/gin-gonic/gin"
)

type authHandler struct {
	log         logger.Logger
	group       *gin.RouterGroup
	authUsecase usecases.Auth
}

func NewAuthHandler(log logger.Logger, group *gin.RouterGroup, authUsecase usecases.Auth) *authHandler {
	return &authHandler{
		log:         log,
		group:       group,
		authUsecase: authUsecase,
	}
}

func (ah *authHandler) InitRoutes() {
	ah.group.POST("/sign_up", ah.signUp)
	ah.group.POST("/sign_in", ah.signIn)
	ah.group.POST("/sign_out", ah.signOut)
	ah.group.POST("/refresh", ah.refresh)
}

func (ah *authHandler) signUp(c *gin.Context) {
	var userDTO *dto.UserDTO
	if err := c.BindJSON(&userDTO); err != nil {
		newHttpErrorResponse(c, ah.log, http.StatusBadRequest, err)
		return
	}

	if err := ah.authUsecase.SignUp(c.Request.Context(), userDTO); err != nil {
		var errAlreadyExists *errs.ErrAlreadyExists
		if errors.As(err, &errAlreadyExists) {
			newHttpErrorResponse(c, ah.log, http.StatusBadRequest, err)
			return
		}

		newHttpErrorResponse(c, ah.log, http.StatusInternalServerError, err)
		return
	}

	newHttpSuccessResponse(c, "ok")
}

func (ah *authHandler) signIn(c *gin.Context) {
	var userDTO *dto.UserDTO

	if err := c.BindJSON(&userDTO); err != nil {
		newHttpErrorResponse(c, ah.log, http.StatusBadRequest, err)
		return
	}

	tokens, err := ah.authUsecase.SignIn(c.Request.Context(), userDTO)
	if err != nil {
		status := http.StatusInternalServerError

		var errNotFound *errs.ErrNotFound
		if errors.As(err, &errNotFound) || errors.Is(err, errs.ErrIncorrectPassword) {
			status = http.StatusUnauthorized
		}

		newHttpErrorResponse(c, ah.log, status, err)
		return
	}

	newHttpSuccessResponse(c, tokens)
}

func (ah *authHandler) signOut(c *gin.Context) {

}

func (ah *authHandler) refresh(c *gin.Context) {
	var refreshToken *dto.RefreshTokenDTO
	if err := c.BindJSON(&refreshToken); err != nil {
		newHttpErrorResponse(c, ah.log, http.StatusBadRequest, err)
		return
	}

	tokens, err := ah.authUsecase.RefreshTokens(c.Request.Context(), refreshToken)
	if err != nil {
		status := http.StatusInternalServerError

		var errNotFound *errs.ErrNotFound
		if errors.As(err, &errNotFound) || errors.Is(err, errs.ErrRefreshTokenExpired) {
			status = http.StatusUnauthorized
		}

		newHttpErrorResponse(c, ah.log, status, err)
		return
	}

	newHttpSuccessResponse(c, tokens)
}
