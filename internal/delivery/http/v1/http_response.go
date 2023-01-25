package v1

import (
	"net/http"

	"github.com/Unlites/knowledge_keeper/pkg/logger"
	"github.com/gin-gonic/gin"
)

type httpResponse struct {
	Status string      `json:"status"`
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
}

func newHttpSuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, httpResponse{
		Status: "success",
		Code:   http.StatusOK,
		Data:   data,
	})
}

func newHttpErrorResponse(c *gin.Context, log logger.Logger, code int, err error) {
	log.Error("", err)
	c.AbortWithStatusJSON(code, httpResponse{
		Status: "error",
		Code:   code,
		Data:   err.Error(),
	})
}
