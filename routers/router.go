package routers

import (
	"API/controller"
	middleware "API/middelware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/register", controller.RegHandler)
		api.POST("/login", controller.LoginHandler)
		api.GET("/file", middleware.AuthByToken(), controller.FileHandler)
	}
	return router
}
