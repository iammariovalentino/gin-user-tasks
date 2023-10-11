package usecase

import (
	"context"
	"gin-user-tasks/src/app/tasks/model"
	"gin-user-tasks/src/app/tasks/query"
	"gin-user-tasks/src/app/tasks/schema"
)

type (
	TaskUsecase interface {
		GetAllTasks(ctx context.Context) (*schema.GetAllTasksResponse, error)
		InsertTask(ctx context.Context, req *schema.InsertTaskRequest) (*schema.InsertTaskResponse, error)
		GetTaskByID(ctx context.Context, id int64) (*schema.GetTaskByIDResponse, error)
		EditTaskByID(ctx context.Context, id int64, req *schema.EditTaskRequest) (*schema.EditTaskResponse, error)
		DeleteTaskByID(ctx context.Context, id int64) error
	}

	taskUsecase struct {
		query query.TaskQuery
	}
)

func NewTaskUsecase(query query.TaskQuery) TaskUsecase {
	return &taskUsecase{query: query}
}

func (u *taskUsecase) GetAllTasks(ctx context.Context) (*schema.GetAllTasksResponse, error) {
	result, err := u.query.GetAllTasks(ctx)
	if err != nil {
		return nil, err
	}

	return &schema.GetAllTasksResponse{Tasks: result}, nil
}

func (u *taskUsecase) InsertTask(ctx context.Context, req *schema.InsertTaskRequest) (*schema.InsertTaskResponse, error) {
	result, err := u.query.InsertTask(ctx, &model.Task{
		UserID:      req.UserID,
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
	})

	if err != nil {
		return nil, err
	}

	return &schema.InsertTaskResponse{
		Task: result,
	}, nil
}

func (u *taskUsecase) GetTaskByID(ctx context.Context, id int64) (*schema.GetTaskByIDResponse, error) {
	result, err := u.query.GetTaskByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &schema.GetTaskByIDResponse{User: result}, nil
}

func (u *taskUsecase) EditTaskByID(ctx context.Context, id int64, req *schema.EditTaskRequest) (*schema.EditTaskResponse, error) {
	task := &model.Task{
		UserID:      req.UserID,
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
	}
	result, err := u.query.UpdateTaskByID(ctx, id, task)

	if err != nil {
		return nil, err
	}

	return &schema.EditTaskResponse{Task: result}, nil
}

func (u *taskUsecase) DeleteTaskByID(ctx context.Context, id int64) error {
	return u.query.DeleteTaskByID(ctx, id)
}
