package database

import (
	"context"
	"net/http"
	// "crypto/des"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

var db *pgx.Conn

func InitDB() {
	// Get connection details from environment variables
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5434")
	user := getEnv("DB_USER", "todoadmin")
	password := getEnv("DB_PASSWORD", "password")
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
func GetTodos(db *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context )  {
		if db == nil {
			log.Println("DB is nil")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "database not initialized"})
			return
		}
		
		rows, err := db.Query(context.Background(), "SELECT id, task, description, completed FROM todos")

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch todos: " + err.Error()})
		return
	}

	defer rows.Close()

	var todos []map[string]interface{}
	for rows.Next() {
		var id int
		var task string
		var description string
		var completed bool
		err := rows.Scan(&id, &task, &description, &completed)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error scanning row: " + err.Error()})
			return
		}
		todos = append(todos, gin.H{"id": id, "task": task})
	}
	c.JSON(200, todos)
	}
}

// add a new todo
func AddTodo(db *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
	task := c.PostForm("task")
	if task == "" {
		c.JSON(400, gin.H{"error": "Task cannot be empty"})
		return
	}

	description := c.PostForm("description")
	if description == "" {
		c.JSON(400, gin.H{"error": "Description cannot be empty"})
		return
	}
	completed := c.PostForm("completed")
	if completed == "" {
		c.JSON(400, gin.H{"error": "Completed cannot be empty"})
		return
	}
	
	_, err := db.Exec(context.Background(), "INSERT INTO todos (task , description, completed) VALUES ($1)", task, description, completed)

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to add todo: " + err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Todo added successfully"})
	}
}

// delete a todo
func DeleteTodo(db *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
	id := c.Param("id")
	_, err := db.Exec(context.Background(), "DELETE FROM todos WHERE id = $1", id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete todo: " + err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Todo deleted successfully"})
	}
}