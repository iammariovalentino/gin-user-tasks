package model

import "time"

type (
	User struct {
		ID        int64     `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement"`
		Name      string    `json:"name" gorm:"column:name;type:varchar(255)"`
		Email     string    `json:"email" gorm:"column:email;type:varchar(255);index:email,unique"`
		Password  string    `json:"-" gorm:"column:password;type:varchar(255)"`
		CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
		UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
	}
)
