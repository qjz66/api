package controller

import (
	"API/model"
	"API/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegHandler(c *gin.Context) {
	svc := service.NewService(c)
	user := model.User{}
	user.Username = c.PostForm("username")
	user.Password = c.PostForm("password")
	id, err := svc.UserReg(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User Regist success",
		"id":      id,
	})
}
