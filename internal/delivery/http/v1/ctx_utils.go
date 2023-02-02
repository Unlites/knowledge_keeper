package v1

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *v1Handler) getUserId(c *gin.Context) (uint, error) {
	userId, err := strconv.ParseUint(c.GetString("userId"), 10, 64)
	if err != nil {
		return 0, err
	}

	return uint(userId), nil
}

func (h *v1Handler) getIdParam(c *gin.Context) (uint, error) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return 0, err
	}

	return uint(id), nil
}
