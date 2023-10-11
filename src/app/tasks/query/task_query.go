package query

import (
	"context"
	"gin-user-tasks/src/app/tasks/model"

	"gorm.io/gorm"
)

type (
	TaskQuery interface {
		GetAllTasks(ctx context.Context) ([]*model.Task, error)
		InsertTask(ctx context.Context, task *model.Task) (*model.Task, error)
		GetTaskByID(ctx context.Context, id int64) (*model.Task, error)
		UpdateTaskByID(ctx context.Context, id int64, task *model.Task) (*model.Task, error)
		DeleteTaskByID(ctx context.Context, id int64) error
	}

	taskQuery struct {
		db *gorm.DB
	}
)

func NewTaskQuery(db *gorm.DB) TaskQuery {
	return &taskQuery{db: db}
}

func (q *taskQuery) GetAllTasks(ctx context.Context) ([]*model.Task, error) {
	tasks := []*model.Task{}

	err := q.db.Find(&tasks).Error
	if err != nil {
		return []*model.Task{}, nil
	}

	return tasks, nil
}

func (q *taskQuery) InsertTask(ctx context.Context, task *model.Task) (*model.Task, error) {
	create := q.db.Create(task)
	if create.Error != nil {
		return nil, create.Error
	}

	return task, nil
}

func (q *taskQuery) GetTaskByID(ctx context.Context, id int64) (*model.Task, error) {
	task := model.Task{}
	err := q.db.Where("id=?", id).First(&task).Error
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (q *taskQuery) UpdateTaskByID(ctx context.Context, id int64, task *model.Task) (*model.Task, error) {
	t := model.Task{}
	err := q.db.Where("id=?", id).First(&t).Error
	if err != nil {
		return nil, err
	}

	err = q.db.Where("id=?", id).Updates(task).Error
	if err != nil {
		return nil, err
	}

	task.ID = id
	task.CreatedAt = t.CreatedAt
	return task, nil
}

func (q *taskQuery) DeleteTaskByID(ctx context.Context, id int64) error {
	task := model.Task{}

	err := q.db.Where("id=?", id).First(&task).Error
	if err != nil {
		return err
	}

	return q.db.Where("id=?", id).Delete(&task).Error
}
