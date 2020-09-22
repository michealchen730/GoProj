package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine{
	engine := gin.Default()

	router := engine.Group("/")
	{
		router.GET("test", func(context *gin.Context) {
			context.String(http.StatusOK,"welcome")
		})
	}
	return engine
}