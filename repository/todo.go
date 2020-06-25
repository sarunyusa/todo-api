package repository

import (
	"context"
	"github.com/jinzhu/gorm"
	"net/http"
	"todo/entity"
	pkgerror "todo/pkg/error"
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
	dbResult := db.Create(t)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return t, nil
}

func (r *todoRepository) UpdateTodo(ctx context.Context, db *gorm.DB, t *entity.Todo) (*entity.Todo, error) {
	dbResult := db.Save(t)
	if dbResult.RecordNotFound() {
		return nil, pkgerror.NewHttpError(http.StatusNotFound, dbResult.Error)
	} else if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return t, nil
}

func (r *todoRepository) DeleteTodo(ctx context.Context, db *gorm.DB, id string) error {
	t := &entity.Todo{}
	dbResult := db.Where("id = ?", id).Delete(t)

	if dbResult.RecordNotFound() {
		return pkgerror.NewHttpError(http.StatusNotFound, dbResult.Error)
	}

	return dbResult.Error
}

func (r *todoRepository) GetById(ctx context.Context, db *gorm.DB, id string) (*entity.Todo, error) {
	t := &entity.Todo{}
	dbResult := db.Where("id = ?", id).First(t)
	if dbResult.RecordNotFound() {
		return nil, pkgerror.NewHttpError(http.StatusNotFound, dbResult.Error)
	} else if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return t, nil
}

func (r *todoRepository) GetNotDone(ctx context.Context, db *gorm.DB) (*[]entity.Todo, error) {
	t := make([]entity.Todo, 0)
	dbResult := db.Where("is_done = ?", false).Find(&t)
	if dbResult.RecordNotFound() {
		return nil, pkgerror.NewHttpError(http.StatusNotFound, dbResult.Error)
	} else if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return &t, nil
}

func NewTodoRepository() TodoRepository {
	return &todoRepository{}
}
