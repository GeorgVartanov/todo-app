package main

import (
	"github.com/gin-gonic/gin"
	"github.ru/GeorgVartanov/todo-app/pkg/database/postgres"
	"github.ru/GeorgVartanov/todo-app/pkg/user/controllers"
	"github.ru/GeorgVartanov/todo-app/pkg/user/routes"
	"github.ru/GeorgVartanov/todo-app/pkg/user/services"
	"github.ru/GeorgVartanov/todo-app/pkg/user/storage/pg"
	"log"
)

func main() {
	db := postgres.NewDB()
	defer db.Close()
	router := gin.Default()
	API := router.Group("/api/")
	userRouter := API.Group("/user/")

	userStorage, err := pg.NewStorage(db)
	if err != nil {
		log.Fatal(err)
	}
	userService := services.NewService(userStorage)
	userControllers := controllers.NewControllers(userService)
	routes.Handler(userRouter, userControllers)

	router.Run()
}
