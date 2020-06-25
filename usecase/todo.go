package usecase

import (
	"context"
	"github.com/jinzhu/gorm"
	"net/http"
	"todo/entity"
	pkgcontext "todo/pkg/context"
	pkgerror "todo/pkg/error"
	"todo/pkg/model"
	"todo/pkg/stopwatch"
	"todo/pkg/util"
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

func (u *todoUseCase) CreateTodo(ctx context.Context, content *model.TodoContent) (*model.TodoInfo, error) {
	log := pkgcontext.GetLoggerFromContext(ctx).WithServiceInfo("CreateTodo")
	l := log.WithField("content", util.Stringify(content))
	defer stopwatch.StartWithLogger(l).Stop()

	if err := content.Validate(); err != nil {
		log.WithError(err).Error("validate content error")
		httpErr := pkgerror.NewHttpError(http.StatusBadRequest, err)
		return nil, httpErr
	}

	t := &entity.Todo{
		Topic:   content.Topic,
		Detail:  content.Detail,
		DueDate: content.DueDate,
		IsDone:  false,
	}

	result, err := func(ctx context.Context, db *gorm.DB) (*entity.Todo, error) {

		defer db.RollbackUnlessCommitted()

		res, err := u.todoRepo.CreateTodo(ctx, db, t)
		if err != nil {
			log.WithError(err).Error("create todo error")
			return nil, err
		}

		if err = db.Commit().Error; err != nil {
			log.WithError(err).Error("commit error")
			return nil, err
		}

		return res, nil
	}(ctx, u.db.Begin())
	if err != nil {
		log.WithError(err).Error("create error")
		return nil, err
	}

	res := &model.TodoInfo{
		TodoContent: model.TodoContent{
			Topic:   result.Topic,
			Detail:  result.Detail,
			DueDate: result.DueDate,
		},
		ID:       result.ID,
		IsDone:   result.IsDone,
		CreateAt: result.CreatedAt,
		UpdateAt: result.UpdatedAt,
	}

	return res, nil
}

func (u *todoUseCase) UpdateTodo(ctx context.Context, content *model.TodoContent) (*model.TodoInfo, error) {
	panic("implement me")
}

func (u *todoUseCase) DeleteTodo(ctx context.Context, id string) error {
	panic("implement me")
}

func (u *todoUseCase) SetTodoDone(ctx context.Context, id string) error {
	panic("implement me")
}

func (u *todoUseCase) GetNotDoneTodo(ctx context.Context) (*[]model.TodoInfo, error) {
	panic("implement me")
}

func (u *todoUseCase) GetTodoById(ctx context.Context, id string) (*model.TodoInfo, error) {
	panic("implement me")
}

func NewTodoUseCase(db *gorm.DB, todoRepo repository.TodoRepository) TodoUseCase {
	return &todoUseCase{
		db:       db,
		todoRepo: todoRepo,
	}
}
