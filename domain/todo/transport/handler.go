package transport

import (
	"go-server-exmaple/domain/todo/service"
	"go-server-exmaple/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type handler struct {
	todoService service.TodoService
}

func NewHandler(todoService service.TodoService) *handler {
	return &handler{
		todoService: todoService,
	}
}

func (h *handler) CreateTodo(c echo.Context) error {
	ctx := c.Request().Context()

	var todo model.Todo
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := c.Validate(todo); err != nil {
		return err
	}

	id, err := h.todoService.CreateTodo(ctx, todo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// annonymous struct
	result := struct {
		ID int `json:"id"`
	}{
		ID: id,
	}

	return c.JSON(http.StatusCreated, result)
}

func (h *handler) GetTodo(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	todo, err := h.todoService.GetTodo(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, todo)
}

func (h *handler) GetTodos(c echo.Context) error {
	ctx := c.Request().Context()

	todos, err := h.todoService.GetTodos(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, todos)
}
