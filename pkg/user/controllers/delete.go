package controllers

import (
	"github.com/gin-gonic/gin"
	"github.ru/GeorgVartanov/todo-app/pkg/user/models"
	"net/http"
)

func (controller Controllers) Delete() gin.HandlerFunc {
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
		err := controller.serv.Delete(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": "Deleted"})
		return
	}
}
