package main

import (
	"log"

	"task_user_RestAPI/internal/database"
	"task_user_RestAPI/internal/handlers/taskHandlers"
	"task_user_RestAPI/internal/handlers/userHandlers"
	"task_user_RestAPI/internal/repository"
	"task_user_RestAPI/internal/services/taskService"
	"task_user_RestAPI/internal/services/userService"
	"task_user_RestAPI/internal/web/tasks"
	"task_user_RestAPI/internal/web/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database.InitDB()

	// Repositories
	taskRepo := repository.NewTaskRepository(database.DB)
	userRepo := repository.NewUserRepository(database.DB)

	// Services
	taskService := taskService.NewService(taskRepo, userRepo)
	userService := userService.NewService(userRepo, taskService)

	// Handlers
	taskHandler := taskHandlers.NewHandler(taskService)
	userHandler := userHandlers.NewHandler(userService)

	// Echo setup
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	tskStrictHandler := tasks.NewStrictHandler(taskHandler, nil)
	tasks.RegisterHandlers(e, tskStrictHandler)

	usrStrictHandler := users.NewStrictHandler(userHandler, nil)
	users.RegisterHandlers(e, usrStrictHandler)

	log.Println("Starting server at :8080")
	e.Logger.Fatal(e.Start(":8080"))
}
