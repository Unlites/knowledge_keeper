package v1

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func (h *v1Handler) addLoggerRequestId(c *gin.Context) {
	requestId := requestid.Get(c)
	h.log.AddKey("request_id", requestId)
}

func (h *v1Handler) addLoggerUserId(c *gin.Context) {
	userId := "test"
	h.log.AddKey("user_id", userId)
}

func (h *v1Handler) userIdentification(c *gin.Context) {

}
