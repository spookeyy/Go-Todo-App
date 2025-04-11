package main

import (
	"context"
	"log"
	"os"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/gin-gonic/gin"
)

var db *pgx.Conn

func initDB() {
	// Get connection details from environment variables
	host := getEnv("DB_HOST", "")
	port := getEnv("DB_PORT", "")
	user := getEnv("DB_USER", "")
	password := getEnv("DB_PASSWORD", "")
	dbname := getEnv("DB_NAME", "todos")
	
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", 
		user, password, host, port, dbname)
	
	log.Printf("Connecting to database with: %s", connStr)
	
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal("DB connection failed: ", err)
	}
	
	// Test the connection
	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatal("DB ping failed: ", err)
	}
	
	log.Println("Successfully connected to database")
	db = conn
	
	// Initialize the database table if it doesn't exist
	_, err = db.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS todos (
			id SERIAL PRIMARY KEY,
			task TEXT NOT NULL
		)
	`)
	if err != nil {
		log.Fatal("Failed to create table: ", err)
	}
}

// Helper function to get environment variables with defaults
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// fetch all todos
func getTodos(c *gin.Context) {
	rows, err := db.Query(context.Background(), "SELECT id, task FROM todos")

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch todos: " + err.Error()})
		return
	}

	defer rows.Close()

	var todos []map[string]interface{}
	for rows.Next() {
		var id int
		var task string
		err := rows.Scan(&id, &task)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error scanning row: " + err.Error()})
			return
		}
		todos = append(todos, gin.H{"id": id, "task": task})
	}
	c.JSON(200, todos)
}

// add a new todo
func addTodo(c *gin.Context) {
	task := c.PostForm("task")
	if task == "" {
		c.JSON(400, gin.H{"error": "Task cannot be empty"})
		return
	}
	
	_, err := db.Exec(context.Background(), "INSERT INTO todos (task) VALUES ($1)", task)

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to add todo: " + err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Todo added successfully"})
}

// delete a todo
func deleteTodo(c *gin.Context) {
	id := c.Param("id")
	_, err := db.Exec(context.Background(), "DELETE FROM todos WHERE id = $1", id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete todo: " + err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Todo deleted successfully"})
}