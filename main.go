package main

import (
	"fmt"
	"gin-user-tasks/src/pkg/config"
	"gin-user-tasks/src/pkg/service"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Printf("loading environment variable .env file\n")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file")
	}

	config.InitConfiguration()

	env := config.Env.Environment
	if env == "" {
		log.Fatal("empty environment")
	}
	fmt.Printf("environmet APP_ENV=%s\n", env)

	server := service.New()
	server.Run()
}
