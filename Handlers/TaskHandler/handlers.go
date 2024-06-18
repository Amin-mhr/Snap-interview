package TaskHandler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Message struct {
	Message string `json:"message"`
}

func updateTaskHandler(service TaskServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		task := new(Task)
		if err := c.Bind(task); err != nil {
			return err
		}
		if err := service.Update(task); err != nil {
			return c.JSON(http.StatusBadRequest, Message{Message: "unsuccessful"})
		}
		return c.JSON(http.StatusOK, Message{Message: "successful"})
	}
}

func getListHandler(service TaskServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		tasks, err := service.TaskList()
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, tasks)
	}
}

func createTaskHandler(service TaskServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		task := new(Task)
		if err := c.Bind(task); err != nil {
			return err
		}
		if err := service.Create(task); err != nil {
			return c.JSON(http.StatusBadRequest, Message{Message: "unsuccessful"})
		}
		return c.JSON(http.StatusOK, task)
	}
}

func TaskRoute(server *echo.Echo, service TaskServiceInterface) {
	server.PUT("/task", updateTaskHandler(service))
	server.GET("/task", getListHandler(service))
	server.POST("/task", createTaskHandler(service))
}
