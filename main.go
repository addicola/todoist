package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/addicola/todoist/todo"
)

func main() {
	// Init database connection
	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&todo.Todo{})

	// Initialize the Todo handler
	todoHandler := todo.NewTodoHandler(db)

	r := gin.Default()

	r.POST("/todos", todoHandler.CreateTodo)
	
	r.Run(":8080")
}
