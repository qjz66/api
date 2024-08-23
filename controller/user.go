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
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User Register success",
		"id":      id,
	})
}

func LoginHandler(c *gin.Context) {
	svc := service.NewService(c)
	user := model.User{}
	user.Username = c.PostForm("username")
	user.Password = c.PostForm("password")
	id, token, err := svc.Login(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User Login success",
		"id":      id,
		"token":   token,
	})
}
