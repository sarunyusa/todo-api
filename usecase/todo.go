package usecase

import (
	"github.com/jinzhu/gorm"
	"todo/repository"
)

type TodoUseCase interface {
}

type todoUseCase struct {
	db       *gorm.DB
	todoRepo repository.TodoRepository
}

func NewTodoUseCase(db *gorm.DB, todoRepo repository.TodoRepository) TodoUseCase {
	return &todoUseCase{
		db:       db,
		todoRepo: todoRepo,
	}
}
