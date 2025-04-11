package main

import (
	"log"
	// "net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	initDB()

	//Todo routes
	route.GET("/todos", getTodos)
	route.POST("/todos", addTodo)
	route.DELETE("/todos/:id", deleteTodo)


	if err := route.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}