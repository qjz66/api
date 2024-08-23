package routers

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Static("/videos", "./videos")
	api := router.Group("/api")
	{
		api.POST("/register")
		api.POST("/login")
		api.POST("/file")
	}
	return router
}
