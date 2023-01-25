package http

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func (h *Handler) addLoggerRequestId(c *gin.Context) {
	requestId := requestid.Get(c)
	h.log.AddKey("request_id", requestId)
}

func (h *Handler) addLoggerUserId(c *gin.Context) {
	userId := "test"
	h.log.AddKey("user_id", userId)
}
