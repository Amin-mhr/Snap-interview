package main

import (
	"SnapGrocery/Handlers/TaskHandler"
	"github.com/labstack/echo/v4"
)

func main() {
	server := echo.New()
	repo := TaskHandler.NewTaskRepository()
	service := TaskHandler.NewTaskService(repo)
	TaskHandler.TaskRoute(server, service)
	server.Start(":6060")
}
