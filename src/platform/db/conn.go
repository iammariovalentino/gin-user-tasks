package db

import (
	taskModel "gin-user-tasks/src/app/tasks/model"
	userModel "gin-user-tasks/src/app/users/model"
	"log"

	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectGorm(db *sqlx.DB) *gorm.DB {
	gorm, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})

	if err != nil {
		log.Fatalf("error happen when connect to database:%s\n", err)
	}

	gorm.AutoMigrate(&userModel.User{}, &taskModel.Task{})
	return gorm
}
