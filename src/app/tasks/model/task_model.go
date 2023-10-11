package model

import "time"

type (
	Task struct {
		ID          int64     `json:"id" gorm:"column:id;type:int;primaryKey;autoIncrement"`
		UserID      int64     `json:"user_id" gorm:"column:user_id;type:int;foreignKey:UserID"`
		Title       string    `json:"title" gorm:"column:title;type:varchar(255)"`
		Description string    `json:"description" gorm:"column:description;type:text"`
		Status      string    `json:"status" gorm:"column:status;type:varchar(50);default:pending"`
		CreatedAt   time.Time `json:"created_at" gorm:"column:created_at"`
		UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at"`
	}
)
