package routers

import (
	"GoProj1.0/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	engine := gin.Default()

	router := engine.Group("/")
	{
		router.GET("test", func(context *gin.Context) {
			context.String(http.StatusOK, "welcome")
		})

		userRouter := router.Group("user")
		{
			userRouter.POST("", handlers.UserSaveService)
		}

	}
	return engine
}
