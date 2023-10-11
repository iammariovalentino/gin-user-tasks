package schema

import "gin-user-tasks/src/app/users/model"

type (
	RegisterUserRequest struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	RegisterUserResponse struct {
		User  *model.User `json:"user"`
		Token interface{} `json:"token"`
	}

	EditUserURI struct {
		ID int64 `uri:"id" binding:"required,numeric"`
	}

	EditUserRequest struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	EditUserResponse struct {
		User *model.User `json:"user"`
	}

	GetAllUsersResponse struct {
		Users []*model.User `json:"users"`
	}

	GetUserByIDRequest struct {
		ID int64 `uri:"id" binding:"required,numeric"`
	}

	GetUserByIDResponse struct {
		User *model.User `json:"user"`
	}

	DeleteUserByIDRequest struct {
		ID int64 `uri:"id" binding:"required,numeric"`
	}
)
