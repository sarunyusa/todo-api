package repository

import (
	"context"
	"github.com/jinzhu/gorm"
	"todo/entity"
)

type TodoRepository interface {
	CreateTodo(ctx context.Context, db *gorm.DB, t *entity.Todo) (*entity.Todo, error)
	UpdateTodo(ctx context.Context, db *gorm.DB, t *entity.Todo) (*entity.Todo, error)
	DeleteTodo(ctx context.Context, db *gorm.DB, id string) error
	GetById(ctx context.Context, db *gorm.DB, id string) (*entity.Todo, error)
	GetNotDone(ctx context.Context, db *gorm.DB) (*[]entity.Todo, error)
}

type todoRepository struct {
}

func (r *todoRepository) CreateTodo(ctx context.Context, db *gorm.DB, t *entity.Todo) (*entity.Todo, error) {
	panic("implement me")
}

func (r *todoRepository) UpdateTodo(ctx context.Context, db *gorm.DB, t *entity.Todo) (*entity.Todo, error) {
	panic("implement me")
}

func (r *todoRepository) DeleteTodo(ctx context.Context, db *gorm.DB, id string) error {
	panic("implement me")
}

func (r *todoRepository) GetById(ctx context.Context, db *gorm.DB, id string) (*entity.Todo, error) {
	panic("implement me")
}

func (r *todoRepository) GetNotDone(ctx context.Context, db *gorm.DB) (*[]entity.Todo, error) {
	panic("implement me")
}

func NewTodoRepository() TodoRepository {
	return &todoRepository{}
}
