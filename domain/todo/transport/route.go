package transport

import "github.com/labstack/echo/v4"

func (h *handler) Route(g *echo.Group) {
	// https://restfulapi.net/resource-naming/
	g.GET("/todos", h.GetTodos)
	g.GET("/todos/:id", h.GetTodo)
	g.PUT("/todos", h.CreateTodo)
}
