package usecase

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gopkg.in/DATA-DOG/go-sqlmock.v2"
	"net/http"
	"testing"
	"time"
	"todo/entity"
	pkgerror "todo/pkg/error"
	"todo/pkg/model"
	"todo/pkg/ptr"
	"todo/pkg/util"
	"todo/test/database"
	mockRepo "todo/test/mocks/repository"
)

func createMock() (*gorm.DB, sqlmock.Sqlmock, mockRepo.TodoRepository) {
	db, smock := database.NewSqlMockConnection()
	return db, smock, mockRepo.TodoRepository{}
}

func TestTodoUseCase_CreateTodo(t *testing.T) {
	t.Run("create success", func(t *testing.T) {
		db, smock, repo := createMock()

		todo := &entity.Todo{
			Base: entity.Base{
				ID:        util.NewID(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: nil,
			},
			Topic:   "test topic",
			Detail:  ptr.String("test detail"),
			DueDate: ptr.Time(time.Now()),
			IsDone:  false,
		}

		repo.On("CreateTodo", mock.Anything, mock.Anything, mock.Anything).
			Return(todo, nil)

		smock.ExpectBegin()
		smock.ExpectCommit()

		u := NewTodoUseCase(db, &repo)

		res, err := u.CreateTodo(context.Background(), &model.TodoContent{Topic: "test topic"})

		require.Nil(t, err)
		require.NotNil(t, res)

		assert.Equal(t, res.ID, todo.ID)
		assert.Equal(t, res.CreateAt, todo.CreatedAt)
		assert.Equal(t, res.UpdateAt, todo.UpdatedAt)
		assert.Equal(t, res.Topic, todo.Topic)
		assert.Equal(t, res.Detail, todo.Detail)
		assert.Equal(t, res.DueDate, todo.DueDate)
		assert.Equal(t, res.IsDone, todo.IsDone)
	})

	t.Run("create fail, topic is blank", func(t *testing.T) {
		db, _, repo := createMock()

		content := &model.TodoContent{
			Topic:   "",
			Detail:  ptr.String("test detail"),
			DueDate: ptr.Time(time.Now()),
		}

		u := NewTodoUseCase(db, &repo)

		res, err := u.CreateTodo(context.Background(), content)

		require.Nil(t, res)
		require.NotNil(t, err)
		require.True(t, pkgerror.IsHttpError(err))

		httpErr := err.(pkgerror.HttpError)
		assert.Equal(t, httpErr.Code(), http.StatusBadRequest)
		assert.Equal(t, httpErr.Error(), "topic is blank")
	})

	t.Run("crate fail, database error", func(t *testing.T) {
		db, smock, repo := createMock()

		repo.On("CreateTodo", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, fmt.Errorf("db error"))

		smock.ExpectBegin()
		smock.ExpectRollback()

		u := NewTodoUseCase(db, &repo)

		res, err := u.CreateTodo(context.Background(), &model.TodoContent{Topic: "test topic"})

		require.Nil(t, res)
		require.NotNil(t, err)

		assert.Equal(t, err.Error(), "db error")
	})
}
