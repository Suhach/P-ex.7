package main

import (
	"log"

	"P_ex.7/internal/database"
	"P_ex.7/internal/handlers/taskHandlers"
	"P_ex.7/internal/handlers/userHandlers"
	"P_ex.7/internal/services/taskService"
	"P_ex.7/internal/services/userService"
	"P_ex.7/internal/web/tasks"
	"P_ex.7/internal/web/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database.InitDB()
	//task
	taskrepo := taskService.NewTaskRepository(database.DB)
	taskservice := taskService.NewService(taskrepo)
	taskhandler := taskHandlers.NewHandler(taskservice)
	//user
	userrepo := userService.NewTaskRepository(database.DB)
	userservice := userService.NewService(userrepo)
	userhandler := userHandlers.NewHandler(userservice)
	//
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//task
	tskstrictHandler := tasks.NewStrictHandler(taskhandler, nil)
	tasks.RegisterHandlers(e, tskstrictHandler)
	//user
	usrstrictHandler := users.NewStrictHandler(userhandler, nil)
	users.RegisterHandlers(e, usrstrictHandler)
	//
	log.Println("Starting server at :8080")
	e.Logger.Fatal(e.Start(":8080"))
}
