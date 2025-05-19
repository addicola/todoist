package todo

import (
	"gorm.io/gorm"
)

type Todo struct {
	Title string `json:"title"`
	Description string `json:"description"`
}

type TodoHandler struct {
	db *gorm.DB
}
