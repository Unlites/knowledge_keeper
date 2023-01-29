package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type httpResponse struct {
	Status string      `json:"status"`
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
}

func (h *v1Handler) newHttpSuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, httpResponse{
		Status: "success",
		Code:   http.StatusOK,
		Data:   data,
	})
}

func (h *v1Handler) newHttpErrorResponse(c *gin.Context, code int, err error) {
	h.log.Error("", err)
	c.AbortWithStatusJSON(code, httpResponse{
		Status: "error",
		Code:   code,
		Data:   err.Error(),
	})
}
