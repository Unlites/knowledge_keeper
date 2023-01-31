package v1

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func (h *v1Handler) addLoggerRequestId(c *gin.Context) {
	requestId := requestid.Get(c)
	h.log.AddKey("request_id", requestId)
}

func (h *v1Handler) addLoggerUserId(c *gin.Context) {
	h.log.AddKey("user_id", c.GetString("userId"))
}

func (h *v1Handler) userIdentification(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		h.newHttpErrorResponse(c, http.StatusUnauthorized, errors.New("empty authorization header"))
		return
	}

	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
		h.newHttpErrorResponse(c, http.StatusUnauthorized, errors.New("invalid authorization header"))
		return
	}

	if len(authHeaderParts[1]) == 0 {
		h.newHttpErrorResponse(c, http.StatusUnauthorized, errors.New("access token is empty"))
		return
	}

	userId, err := h.usecases.Auth.ParseUserIdFromAccessToken(c.Request.Context(), authHeaderParts[1])
	if err != nil {
		h.newHttpErrorResponse(c, http.StatusUnauthorized, err)
		return
	}

	c.Set("userId", userId)
}
