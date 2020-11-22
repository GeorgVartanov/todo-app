package routes

import (
	"github.com/gin-gonic/gin"
	"github.ru/GeorgVartanov/todo-app/pkg/user/controllers"
)

func Handler(router *gin.RouterGroup, controllers controllers.UserControllers) {
	//router.Use(controllers.Checker())
	router.POST("/create/", controllers.Create())
	router.GET("/all/", controllers.ReadAll())
	router.POST("/one/", controllers.Read())
	router.DELETE("/delete/", controllers.Delete())
	router.PUT("/update/", controllers.Update())
}
