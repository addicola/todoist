package todo

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	// Initialize the database connection
	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Todo{})

	return db
}

func setupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	todoHandler := NewTodoHandler(db)

	r.GET("/todos", todoHandler.GetTodos)
	r.POST("/todos", todoHandler.CreateTodo)
	r.PATCH("/todos/:id", todoHandler.UpdateTodo)
	r.DELETE("/todos/:id", todoHandler.DeleteTodo)

	return r
}

func TestCreateTodo_Success(t *testing.T) {
	db := setupTestDB()
	defer db.Exec("DROP TABLE todos") // Clean up the database after the test

	// Setup the router
	router := setupRouter(db)

	// Create a new todo
	todo := Todo{Title: "Test Todo", Description: "This is a test todo"}
	body, _ := json.Marshal(todo)

	req, _ := http.NewRequest("POST", "/todos", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var createdTodo Todo
	err := json.Unmarshal(w.Body.Bytes(), &createdTodo)
	assert.NoError(t, err)
	assert.Equal(t, "Test Todo", createdTodo.Title)
	assert.Equal(t, "This is a test todo", createdTodo.Description)
}

func TestCreateTodo_Failure(t *testing.T) {
	db := setupTestDB()
	defer db.Exec("DROP TABLE todos") // Clean up the database after the test

	// Setup the router
	router := setupRouter(db)

	// Create a new todo with missing title
	todo := Todo{Description: "This is a test todo"}
	body, _ := json.Marshal(todo)

	req, _ := http.NewRequest("POST", "/todos", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
