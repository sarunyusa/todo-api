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

func TestTodoUseCase_UpdateTodo(t *testing.T) {
	t.Run("update success", func(t *testing.T) {
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

		repo.On("GetById", mock.Anything, mock.Anything, mock.Anything).
			Return(todo, nil)

		repo.On("UpdateTodo", mock.Anything, mock.Anything, mock.Anything).
			Return(todo, nil)

		smock.ExpectBegin()
		smock.ExpectCommit()

		u := NewTodoUseCase(db, &repo)

		res, err := u.UpdateTodo(context.Background(), todo.ID, &model.TodoContent{Topic: "test topic"})

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

	t.Run("update fail, id is blank", func(t *testing.T) {
		db, _, repo := createMock()

		content := &model.TodoContent{
			Topic:   "",
			Detail:  ptr.String("test detail"),
			DueDate: ptr.Time(time.Now()),
		}

		u := NewTodoUseCase(db, &repo)

		res, err := u.UpdateTodo(context.Background(), "", content)

		require.Nil(t, res)
		require.NotNil(t, err)
		require.True(t, pkgerror.IsHttpError(err))

		httpErr := err.(pkgerror.HttpError)
		assert.Equal(t, httpErr.Code(), http.StatusBadRequest)
		assert.Equal(t, httpErr.Error(), "id is blank")
	})

	t.Run("update fail, topic is blank", func(t *testing.T) {
		db, _, repo := createMock()

		content := &model.TodoContent{
			Topic:   "",
			Detail:  ptr.String("test detail"),
			DueDate: ptr.Time(time.Now()),
		}

		u := NewTodoUseCase(db, &repo)

		res, err := u.UpdateTodo(context.Background(), util.NewID(), content)

		require.Nil(t, res)
		require.NotNil(t, err)
		require.True(t, pkgerror.IsHttpError(err))

		httpErr := err.(pkgerror.HttpError)
		assert.Equal(t, httpErr.Code(), http.StatusBadRequest)
		assert.Equal(t, httpErr.Error(), "topic is blank")
	})

	t.Run("update fail, get todo error", func(t *testing.T) {
		db, _, repo := createMock()

		repo.On("GetById", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, fmt.Errorf("get todo error"))

		u := NewTodoUseCase(db, &repo)

		res, err := u.UpdateTodo(context.Background(), util.NewID(), &model.TodoContent{Topic: "test topic"})

		require.Nil(t, res)
		require.NotNil(t, err)

		assert.Equal(t, err.Error(), "get todo error")
	})

	t.Run("update fail, database error", func(t *testing.T) {
		db, smock, repo := createMock()

		repo.On("GetById", mock.Anything, mock.Anything, mock.Anything).
			Return(&entity.Todo{}, nil)

		repo.On("UpdateTodo", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, fmt.Errorf("db error"))

		smock.ExpectBegin()
		smock.ExpectRollback()

		u := NewTodoUseCase(db, &repo)

		res, err := u.UpdateTodo(context.Background(), util.NewID(), &model.TodoContent{Topic: "test topic"})

		require.Nil(t, res)
		require.NotNil(t, err)

		assert.Equal(t, err.Error(), "db error")
	})
}

func TestTodoUseCase_DeleteTodo(t *testing.T) {
	t.Run("delete success", func(t *testing.T) {
		db, smock, repo := createMock()

		id := util.NewID()

		repo.On("DeleteTodo", mock.Anything, mock.Anything, mock.Anything).
			Return(nil)

		smock.ExpectBegin()
		smock.ExpectCommit()

		u := NewTodoUseCase(db, &repo)

		err := u.DeleteTodo(context.Background(), id)

		require.Nil(t, err)
	})

	t.Run("delete fail, id is blank", func(t *testing.T) {
		db, _, repo := createMock()

		u := NewTodoUseCase(db, &repo)

		err := u.DeleteTodo(context.Background(), "")

		require.NotNil(t, err)
		require.True(t, pkgerror.IsHttpError(err))

		httpErr := err.(pkgerror.HttpError)
		assert.Equal(t, httpErr.Code(), http.StatusBadRequest)
		assert.Equal(t, httpErr.Error(), "id is blank")
	})

	t.Run("delete fail, database error", func(t *testing.T) {
		db, smock, repo := createMock()

		id := util.NewID()

		repo.On("DeleteTodo", mock.Anything, mock.Anything, mock.Anything).
			Return(fmt.Errorf("db error"))

		smock.ExpectBegin()
		smock.ExpectRollback()

		u := NewTodoUseCase(db, &repo)

		err := u.DeleteTodo(context.Background(), id)

		require.NotNil(t, err)

		assert.Equal(t, err.Error(), "db error")
	})
}

func TestTodoUseCase_SetTodoDone(t *testing.T) {
	t.Run("set done success", func(t *testing.T) {
		db, smock, repo := createMock()

		id := util.NewID()

		repo.On("GetById", mock.Anything, mock.Anything, mock.Anything).
			Return(&entity.Todo{}, nil)

		repo.On("UpdateTodo", mock.Anything, mock.Anything, mock.Anything).
			Return(&entity.Todo{}, nil)

		smock.ExpectBegin()
		smock.ExpectCommit()

		u := NewTodoUseCase(db, &repo)

		err := u.SetTodoDone(context.Background(), id)

		require.Nil(t, err)
	})

	t.Run("set done fail, id is blank", func(t *testing.T) {
		db, _, repo := createMock()

		u := NewTodoUseCase(db, &repo)

		err := u.SetTodoDone(context.Background(), "")

		require.NotNil(t, err)
		require.True(t, pkgerror.IsHttpError(err))

		httpErr := err.(pkgerror.HttpError)
		assert.Equal(t, httpErr.Code(), http.StatusBadRequest)
		assert.Equal(t, httpErr.Error(), "id is blank")
	})

	t.Run("set done fail, get todo error", func(t *testing.T) {
		db, _, repo := createMock()

		id := util.NewID()

		repo.On("GetById", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, fmt.Errorf("get todo error"))

		u := NewTodoUseCase(db, &repo)

		err := u.SetTodoDone(context.Background(), id)

		require.NotNil(t, err)

		assert.Equal(t, err.Error(), "get todo error")
	})

	t.Run("set done fail, database error", func(t *testing.T) {
		db, smock, repo := createMock()

		id := util.NewID()

		repo.On("GetById", mock.Anything, mock.Anything, mock.Anything).
			Return(&entity.Todo{}, nil)

		repo.On("UpdateTodo", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, fmt.Errorf("db error"))

		smock.ExpectBegin()
		smock.ExpectRollback()

		u := NewTodoUseCase(db, &repo)

		err := u.SetTodoDone(context.Background(), id)

		require.NotNil(t, err)

		assert.Equal(t, err.Error(), "db error")
	})
}

func TestTodoUseCase_GetTodoById(t *testing.T) {
	t.Run("get by id success", func(t *testing.T) {
		db, _, repo := createMock()

		id := util.NewID()

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

		repo.On("GetById", mock.Anything, mock.Anything, mock.Anything).
			Return(todo, nil)

		u := NewTodoUseCase(db, &repo)

		res, err := u.GetTodoById(context.Background(), id)

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

	t.Run("get by id fail, id is blank", func(t *testing.T) {
		db, _, repo := createMock()

		u := NewTodoUseCase(db, &repo)

		res, err := u.GetTodoById(context.Background(), "")

		require.Nil(t, res)
		require.NotNil(t, err)
		require.True(t, pkgerror.IsHttpError(err))

		httpErr := err.(pkgerror.HttpError)
		assert.Equal(t, httpErr.Code(), http.StatusBadRequest)
		assert.Equal(t, httpErr.Error(), "id is blank")
	})

	t.Run("set done fail, get todo error", func(t *testing.T) {
		db, _, repo := createMock()

		id := util.NewID()

		repo.On("GetById", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, fmt.Errorf("get todo error"))

		u := NewTodoUseCase(db, &repo)

		res, err := u.GetTodoById(context.Background(), id)

		require.Nil(t, res)
		require.NotNil(t, err)

		assert.Equal(t, err.Error(), "get todo error")
	})
}

func TestTodoUseCase_GetNotDoneTodo(t *testing.T) {
	t.Run("get not done success", func(t *testing.T) {
		db, _, repo := createMock()

		tds := make([]entity.Todo, 3)

		tds[0] = entity.Todo{
			Base: entity.Base{
				ID:        util.NewID(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: nil,
			},
			Topic:   "test topic1",
			Detail:  ptr.String("test detail1"),
			DueDate: ptr.Time(time.Now()),
			IsDone:  false,
		}

		tds[1] = entity.Todo{
			Base: entity.Base{
				ID:        util.NewID(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: nil,
			},
			Topic:   "test topic2",
			Detail:  ptr.String("test detail2"),
			DueDate: ptr.Time(time.Now()),
			IsDone:  false,
		}

		tds[2] = entity.Todo{
			Base: entity.Base{
				ID:        util.NewID(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: nil,
			},
			Topic:   "test topic3",
			Detail:  ptr.String("test detail3"),
			DueDate: ptr.Time(time.Now()),
			IsDone:  false,
		}

		repo.On("GetNotDone", mock.Anything, mock.Anything).
			Return(&tds, nil)

		u := NewTodoUseCase(db, &repo)

		res, err := u.GetNotDoneTodo(context.Background())

		require.Nil(t, err)
		require.NotNil(t, res)

		list := *res
		assert.Equal(t, 3, len(list))

		for i, res := range list {
			assert.Equal(t, res.ID, tds[i].ID)
			assert.Equal(t, res.CreateAt, tds[i].CreatedAt)
			assert.Equal(t, res.UpdateAt, tds[i].UpdatedAt)
			assert.Equal(t, res.Topic, tds[i].Topic)
			assert.Equal(t, res.Detail, tds[i].Detail)
			assert.Equal(t, res.DueDate, tds[i].DueDate)
			assert.Equal(t, res.IsDone, tds[i].IsDone)
		}
	})

	t.Run("set done fail, db error", func(t *testing.T) {
		db, _, repo := createMock()

		repo.On("GetNotDone", mock.Anything, mock.Anything).
			Return(nil, fmt.Errorf("db error"))

		u := NewTodoUseCase(db, &repo)

		res, err := u.GetNotDoneTodo(context.Background())

		require.NotNil(t, err)
		require.Nil(t, res)

		assert.Equal(t, err.Error(), "db error")
	})
}
