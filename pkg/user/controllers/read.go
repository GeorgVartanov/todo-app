package controllers

import (
	"github.com/gin-gonic/gin"
	"github.ru/GeorgVartanov/todo-app/pkg/user/models"
	"net/http"
)

func (controller Controllers) ReadAll() gin.HandlerFunc {
	return func(c *gin.Context) {

		userDB, err := controller.serv.ReadAll()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": &userDB})
		return
	}
}

func (controller Controllers) Read() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.UserIn
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := user.ValidateEmail(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		userDB, err := controller.serv.Read(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": &userDB})
		return
	}
}
