package controllers

import (
	"github.com/gin-gonic/gin"
	"github.ru/GeorgVartanov/todo-app/pkg/user/services"
)

type UserControllers interface {
	Create() gin.HandlerFunc
	ReadAll() gin.HandlerFunc
	Read() gin.HandlerFunc
	Delete() gin.HandlerFunc
	Update() gin.HandlerFunc
}

type Controllers struct {
	serv services.Service
}

func NewControllers(service services.Service) UserControllers {
	return &Controllers{service}
}


