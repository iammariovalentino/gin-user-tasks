package schema

import "gin-user-tasks/src/app/tasks/model"

type (
	InsertTaskRequest struct {
		UserID      int64  `json:"user_id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Status      string `json:"status"`
	}

	InsertTaskResponse struct {
		Task *model.Task `json:"task"`
	}

	EditTaskURI struct {
		ID int64 `uri:"id" binding:"required,numeric"`
	}

	EditTaskRequest struct {
		UserID      int64  `json:"user_id" binding:"required"`
		Title       string `json:"title" binding:"required"`
		Description string `json:"description" binding:"required"`
		Status      string `json:"status" binding:"required"`
	}

	EditTaskResponse struct {
		Task *model.Task `json:"task"`
	}

	GetAllTasksResponse struct {
		Tasks []*model.Task `json:"tasks"`
	}

	GetTaskByIDRequest struct {
		ID int64 `uri:"id" binding:"required,numeric"`
	}

	GetTaskByIDResponse struct {
		User *model.Task `json:"task"`
	}

	DeleteUserByIDRequest struct {
		ID int64 `uri:"id" binding:"required,numeric"`
	}
)
