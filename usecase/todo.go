package usecase

import (
	"context"
	"github.com/jinzhu/gorm"
	"todo/pkg/model"
	"todo/repository"
)

type TodoUseCase interface {
	CreateTodo(ctx context.Context, content *model.TodoContent) (*model.TodoInfo, error)
	UpdateTodo(ctx context.Context, content *model.TodoContent) (*model.TodoInfo, error)
	DeleteTodo(ctx context.Context, id string) error
	SetTodoDone(ctx context.Context, id string) error

	GetNotDoneTodo(ctx context.Context) (*[]model.TodoInfo, error)
	GetTodoById(ctx context.Context, id string) (*model.TodoInfo, error)
}

type todoUseCase struct {
	db       *gorm.DB
	todoRepo repository.TodoRepository
}

func (t *todoUseCase) CreateTodo(ctx context.Context, content *model.TodoContent) (*model.TodoInfo, error) {
	panic("implement me")
}

func (t *todoUseCase) UpdateTodo(ctx context.Context, content *model.TodoContent) (*model.TodoInfo, error) {
	panic("implement me")
}

func (t *todoUseCase) DeleteTodo(ctx context.Context, id string) error {
	panic("implement me")
}

func (t *todoUseCase) SetTodoDone(ctx context.Context, id string) error {
	panic("implement me")
}

func (t *todoUseCase) GetNotDoneTodo(ctx context.Context) (*[]model.TodoInfo, error) {
	panic("implement me")
}

func (t *todoUseCase) GetTodoById(ctx context.Context, id string) (*model.TodoInfo, error) {
	panic("implement me")
}

func NewTodoUseCase(db *gorm.DB, todoRepo repository.TodoRepository) TodoUseCase {
	return &todoUseCase{
		db:       db,
		todoRepo: todoRepo,
	}
}
