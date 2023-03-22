package http

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func NewRouter(origins []string) *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(requestid.New())
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - %s - %s: [%s %s] - [%d | %s]\n",
			param.TimeStamp.Format("2006:01:02 15:04:05"),
			param.ClientIP,
			param.Request.Header.Get("X-Request-Id"),
			param.Method,
			param.Path,
			param.StatusCode,
			param.Latency,
		)
	}))
	router.Use(cors.New(cors.Config{
		AllowOrigins: origins,
	}))

	router.GET("/ping", func(c *gin.Context) { c.String(http.StatusOK, c.GetHeader("X-Request-ID")) })

	return router
}
