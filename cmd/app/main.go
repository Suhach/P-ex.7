package main

import (
	"log"

	"P_ex.7/internal/database"
	"P_ex.7/internal/taskHandlers"
	"P_ex.7/internal/taskService"
	"P_ex.7/internal/web/tasks"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database.InitDB()
	err := database.DB.AutoMigrate(&taskService.Task{})
	if err != nil {
		log.Fatal(err)
	}
	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(repo)
	handler := taskHandlers.NewHandler(service)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(handler, nil)
	tasks.RegisterHandlers(e, strictHandler)

	log.Println("Starting server at :8080")
	e.Logger.Fatal(e.Start(":8080"))
}
