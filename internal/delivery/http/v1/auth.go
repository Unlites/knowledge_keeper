package v1

import (
	"errors"
	"net/http"

	"github.com/Unlites/knowledge_keeper/internal/dto"
	"github.com/Unlites/knowledge_keeper/internal/errs"
	"github.com/gin-gonic/gin"
)

func (h *v1Handler) initAuthRoutes(authGroup *gin.RouterGroup) {
	authGroup.POST("/sign_up", h.signUp)
	authGroup.POST("/sign_in", h.signIn)
	authGroup.POST("/sign_out", h.signOut)
	authGroup.POST("/refresh", h.refresh)
}

func (h *v1Handler) signUp(c *gin.Context) {
	var userDTO *dto.UserDTO
	if err := c.BindJSON(&userDTO); err != nil {
		h.newHttpErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if err := userDTO.Validate(); err != nil {
		h.newHttpErrorResponse(c, http.StatusBadRequest, &errs.ErrValidation{Message: err.Error()})
		return
	}

	if err := h.usecases.Auth.SignUp(c.Request.Context(), userDTO); err != nil {
		var errAlreadyExists *errs.ErrAlreadyExists
		if errors.As(err, &errAlreadyExists) {
			h.newHttpErrorResponse(c, http.StatusBadRequest, err)
			return
		}

		h.newHttpErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	h.newHttpSuccessResponse(c, "ok")
}

func (h *v1Handler) signIn(c *gin.Context) {
	var userDTO *dto.UserDTO

	if err := c.BindJSON(&userDTO); err != nil {
		h.newHttpErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if err := userDTO.Validate(); err != nil {
		h.newHttpErrorResponse(c, http.StatusBadRequest, &errs.ErrValidation{Message: err.Error()})
		return
	}

	tokens, err := h.usecases.Auth.SignIn(c.Request.Context(), userDTO)
	if err != nil {
		status := http.StatusInternalServerError

		var errNotFound *errs.ErrNotFound
		if errors.As(err, &errNotFound) || errors.Is(err, errs.ErrIncorrectPassword) {
			status = http.StatusUnauthorized
		}

		h.newHttpErrorResponse(c, status, err)
		return
	}

	h.newHttpSuccessResponse(c, tokens)
}

func (h *v1Handler) signOut(c *gin.Context) {
	var refreshToken *dto.RefreshTokenDTO
	if err := c.BindJSON(&refreshToken); err != nil {
		h.newHttpErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if err := refreshToken.Validate(); err != nil {
		h.newHttpErrorResponse(c, http.StatusBadRequest, &errs.ErrValidation{Message: err.Error()})
		return
	}

	if err := h.usecases.Auth.SignOut(c.Request.Context(), refreshToken); err != nil {
		h.newHttpErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	h.newHttpSuccessResponse(c, "ok")
}

func (h *v1Handler) refresh(c *gin.Context) {
	var refreshToken *dto.RefreshTokenDTO
	if err := c.BindJSON(&refreshToken); err != nil {
		h.newHttpErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if err := refreshToken.Validate(); err != nil {
		h.newHttpErrorResponse(c, http.StatusBadRequest, &errs.ErrValidation{Message: err.Error()})
		return
	}

	tokens, err := h.usecases.Auth.RefreshTokens(c.Request.Context(), refreshToken)
	if err != nil {
		status := http.StatusInternalServerError

		var errNotFound *errs.ErrNotFound
		if errors.As(err, &errNotFound) || errors.Is(err, errs.ErrRefreshTokenExpired) {
			status = http.StatusUnauthorized
		}

		h.newHttpErrorResponse(c, status, err)
		return
	}

	h.newHttpSuccessResponse(c, tokens)
}
