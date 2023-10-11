package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"gin-user-tasks/src/app/users/model"
	"gin-user-tasks/src/app/users/query"
	"gin-user-tasks/src/app/users/schema"
	"gin-user-tasks/src/pkg/util"
	"io"
	"net/http"
	"time"
)

type (
	UserUsecase interface {
		GetAllUsers(ctx context.Context) (*schema.GetAllUsersResponse, error)
		RegisterUser(ctx context.Context, req *schema.RegisterUserRequest) (*schema.RegisterUserResponse, error)
		GetUserByID(ctx context.Context, id int64) (*schema.GetUserByIDResponse, error)
		EditUserByID(ctx context.Context, id int64, req *schema.EditUserRequest) (*schema.EditUserResponse, error)
		DeleteUserByID(ctx context.Context, id int64) error
	}

	userUsecase struct {
		query query.UserQuery
	}
)

func NewUserUsecase(query query.UserQuery) UserUsecase {
	return &userUsecase{query: query}
}

func (u *userUsecase) GetAllUsers(ctx context.Context) (*schema.GetAllUsersResponse, error) {
	result, err := u.query.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	return &schema.GetAllUsersResponse{Users: result}, nil
}

func (u *userUsecase) RegisterUser(ctx context.Context, req *schema.RegisterUserRequest) (*schema.RegisterUserResponse, error) {
	url := "http://localhost:9096/oauth/access-token"
	r, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	httpClient := &http.Client{
		Transport: &http.Transport{},
		Timeout:   time.Duration(30 * time.Second),
	}

	resp, err := httpClient.Do(r)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(`failed to save user`)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	rs := map[string]interface{}{}
	err = json.Unmarshal(body, &rs)
	if err != nil {
		return nil, err
	}

	result, err := u.query.InsertUser(ctx, &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: util.HashPassword(req.Password),
	})

	if err != nil {
		return nil, err
	}

	var res interface{}
	if _, ok := rs["token"]; ok {
		res = rs["token"]
	}

	return &schema.RegisterUserResponse{
		User:  result,
		Token: res,
	}, nil
}

func (u *userUsecase) GetUserByID(ctx context.Context, id int64) (*schema.GetUserByIDResponse, error) {
	result, err := u.query.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &schema.GetUserByIDResponse{User: result}, nil
}

func (u *userUsecase) EditUserByID(ctx context.Context, id int64, req *schema.EditUserRequest) (*schema.EditUserResponse, error) {
	result, err := u.query.UpdateUserByID(ctx, id, &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: util.HashPassword(req.Password),
	})

	if err != nil {
		return nil, err
	}

	return &schema.EditUserResponse{User: result}, nil
}

func (u *userUsecase) DeleteUserByID(ctx context.Context, id int64) error {
	return u.query.DeleteUserByID(ctx, id)
}
