package usecase

import (
	"context"
	"fmt"
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
	UpdateTodo(ctx context.Context, id string, content *model.TodoContent) (*model.TodoInfo, error)
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

func (u *todoUseCase) UpdateTodo(ctx context.Context, id string, content *model.TodoContent) (*model.TodoInfo, error) {
	log := pkgcontext.GetLoggerFromContext(ctx).WithServiceInfo("UpdateTodo")
	l := log.WithField("id", id).WithField("content", util.Stringify(content))
	defer stopwatch.StartWithLogger(l).Stop()

	if id == "" {
		err := pkgerror.NewHttpError(http.StatusBadRequest, fmt.Errorf("id is blank"))
		log.Error(err)
		return nil, err
	}

	if err := content.Validate(); err != nil {
		log.WithError(err).Error("validate content error")
		httpErr := pkgerror.NewHttpError(http.StatusBadRequest, err)
		return nil, httpErr
	}

	t, err := u.todoRepo.GetById(ctx, u.db, id)
	if err != nil {
		log.WithError(err).Error("get todo error")
		return nil, err
	}

	t.Topic = content.Topic
	t.Detail = content.Detail
	t.DueDate = content.DueDate

	result, err := func(ctx context.Context, db *gorm.DB) (*entity.Todo, error) {

		defer db.RollbackUnlessCommitted()

		res, err := u.todoRepo.UpdateTodo(ctx, db, t)
		if err != nil {
			log.WithError(err).Error("update todo error")
			return nil, err
		}

		if err = db.Commit().Error; err != nil {
			log.WithError(err).Error("commit error")
			return nil, err
		}

		return res, nil
	}(ctx, u.db.Begin())
	if err != nil {
		log.WithError(err).Error("update error")
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

func (u *todoUseCase) DeleteTodo(ctx context.Context, id string) error {
	log := pkgcontext.GetLoggerFromContext(ctx).WithServiceInfo("DeleteTodo")
	l := log.WithField("id", id)
	defer stopwatch.StartWithLogger(l).Stop()

	if id == "" {
		err := pkgerror.NewHttpError(http.StatusBadRequest, fmt.Errorf("id is blank"))
		log.Error(err)
		return err
	}

	err := func(ctx context.Context, db *gorm.DB) error {

		defer db.RollbackUnlessCommitted()

		err := u.todoRepo.DeleteTodo(ctx, db, id)
		if err != nil {
			log.WithError(err).Error("delete todo error")
			return err
		}

		if err = db.Commit().Error; err != nil {
			log.WithError(err).Error("commit error")
			return err
		}

		return nil
	}(ctx, u.db.Begin())
	if err != nil {
		log.WithError(err).Error("delete error")
		return err
	}

	return nil
}

func (u *todoUseCase) SetTodoDone(ctx context.Context, id string) error {
	log := pkgcontext.GetLoggerFromContext(ctx).WithServiceInfo("SetTodoDone")
	l := log.WithField("id", id)
	defer stopwatch.StartWithLogger(l).Stop()

	if id == "" {
		err := pkgerror.NewHttpError(http.StatusBadRequest, fmt.Errorf("id is blank"))
		log.Error(err)
		return err
	}

	t, err := u.todoRepo.GetById(ctx, u.db, id)
	if err != nil {
		log.WithError(err).Error("get todo error")
		return err
	}

	t.IsDone = true

	_, err = func(ctx context.Context, db *gorm.DB) (*entity.Todo, error) {

		defer db.RollbackUnlessCommitted()

		res, err := u.todoRepo.UpdateTodo(ctx, db, t)
		if err != nil {
			log.WithError(err).Error("update todo error")
			return nil, err
		}

		if err = db.Commit().Error; err != nil {
			log.WithError(err).Error("commit error")
			return nil, err
		}

		return res, nil
	}(ctx, u.db.Begin())
	if err != nil {
		log.WithError(err).Error("set done error")
		return err
	}

	return nil
}

func (u *todoUseCase) GetNotDoneTodo(ctx context.Context) (*[]model.TodoInfo, error) {
	log := pkgcontext.GetLoggerFromContext(ctx).WithServiceInfo("GetNotDoneTodo")
	defer stopwatch.StartWithLogger(log).Stop()

	list, err := u.todoRepo.GetNotDone(ctx, u.db)
	if err != nil {
		log.WithError(err).Error("get todo error")
		return nil, err
	}

	res := make([]model.TodoInfo, len(*list))

	for i, td := range *list {
		res[i] = model.TodoInfo{
			TodoContent: model.TodoContent{
				Topic:   td.Topic,
				Detail:  td.Detail,
				DueDate: td.DueDate,
			},
			ID:       td.ID,
			IsDone:   td.IsDone,
			CreateAt: td.CreatedAt,
			UpdateAt: td.UpdatedAt,
		}
	}

	return &res, nil
}

func (u *todoUseCase) GetTodoById(ctx context.Context, id string) (*model.TodoInfo, error) {
	log := pkgcontext.GetLoggerFromContext(ctx).WithServiceInfo("GetTodoById")
	l := log.WithField("id", id)
	defer stopwatch.StartWithLogger(l).Stop()

	if id == "" {
		err := pkgerror.NewHttpError(http.StatusBadRequest, fmt.Errorf("id is blank"))
		log.Error(err)
		return nil, err
	}

	t, err := u.todoRepo.GetById(ctx, u.db, id)
	if err != nil {
		log.WithError(err).Error("get todo error")
		return nil, err
	}

	res := &model.TodoInfo{
		TodoContent: model.TodoContent{
			Topic:   t.Topic,
			Detail:  t.Detail,
			DueDate: t.DueDate,
		},
		ID:       t.ID,
		IsDone:   t.IsDone,
		CreateAt: t.CreatedAt,
		UpdateAt: t.UpdatedAt,
	}

	return res, nil
}

func NewTodoUseCase(db *gorm.DB, todoRepo repository.TodoRepository) TodoUseCase {
	return &todoUseCase{
		db:       db,
		todoRepo: todoRepo,
	}
}
