package query

import (
	"context"
	"errors"
	"gin-user-tasks/src/app/users/model"

	"gorm.io/gorm"
)

type (
	UserQuery interface {
		GetAllUsers(ctx context.Context) ([]*model.User, error)
		InsertUser(ctx context.Context, user *model.User) (*model.User, error)
		GetUserByID(ctx context.Context, id int64) (*model.User, error)
		UpdateUserByID(ctx context.Context, id int64, user *model.User) (*model.User, error)
		DeleteUserByID(ctx context.Context, id int64) error
	}

	userQuery struct {
		db *gorm.DB
	}
)

func NewUserQuery(db *gorm.DB) UserQuery {
	return &userQuery{db: db}
}

func (q *userQuery) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	users := []*model.User{}

	err := q.db.Find(&users).Error
	if err != nil {
		return []*model.User{}, nil
	}

	return users, nil
}

func (q *userQuery) InsertUser(ctx context.Context, user *model.User) (*model.User, error) {
	u := model.User{}
	err := q.db.Where("email=?", user.Email).First(&u).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
		create := q.db.Create(user)
		if create.Error != nil {
			return nil, create.Error
		}
	}

	if user.Email == u.Email {
		err := errors.New(`email is already exist`)
		return nil, err
	}

	return user, nil

}

func (q *userQuery) GetUserByID(ctx context.Context, id int64) (*model.User, error) {
	user := model.User{}
	err := q.db.Where("id=?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (q *userQuery) UpdateUserByID(ctx context.Context, id int64, user *model.User) (*model.User, error) {
	u := model.User{}
	err := q.db.Where("id=?", id).First(&u).Error
	if err != nil {
		return nil, err
	}

	err = q.db.Where("id=?", id).Updates(user).Error
	if err != nil {
		return nil, err
	}

	user.ID = id
	user.CreatedAt = u.CreatedAt
	return user, nil
}

func (q *userQuery) DeleteUserByID(ctx context.Context, id int64) error {
	user := model.User{}

	err := q.db.Where("id=?", id).First(&user).Error
	if err != nil {
		return err
	}

	return q.db.Where("id=?", id).Delete(&user).Error
}
