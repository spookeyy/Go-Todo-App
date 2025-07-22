package main

import (
	"log"
	// "net/http"
	"go-todo-app/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	// _ "github.com/lib/pq"
)

var db *pgx.Conn

func main() {
	route := gin.Default()

	database.InitDB()


	route.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))

	// route.Use(func(c *gin.Context) {
	// 	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// 	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	// 	c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
	// 	if c.Request.Method == "OPTIONS" {
	// 		c.AbortWithStatus(204)
	// 		return
	// 	}
	// 	c.Next()
	// })
	

	//Health check
	route.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Go Todo API Service is up and running"})
	})

	//Todo routes
	route.GET("/todos", database.GetTodos(db))
	route.POST("/todos", database.AddTodo(db))
	route.DELETE("/todos/:id", database.DeleteTodo(db))


	if err := route.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}