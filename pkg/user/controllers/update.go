package controllers

import (
	"github.com/gin-gonic/gin"
	"github.ru/GeorgVartanov/todo-app/pkg/user/models"
	"net/http"
)

func (controller Controllers) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var updateUser models.UserIn
		if err := c.ShouldBindJSON(&updateUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := updateUser.ValidateFieldsUpdate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := controller.serv.Update(&updateUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": "Updated"})
		return
	}
}
