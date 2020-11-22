package controllers

import (
	"github.com/gin-gonic/gin"
	"github.ru/GeorgVartanov/todo-app/pkg/user/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func (controller *Controllers) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newApiUser models.UserIn
		if err := c.ShouldBindJSON(&newApiUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := newApiUser.ValidateFields(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		hashedPasswoed,err:= bcrypt.GenerateFromPassword([]byte(newApiUser.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newApiUser.Password = string(hashedPasswoed)
		err = controller.serv.Create(&newApiUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": "ok"})
		return
	}
}
