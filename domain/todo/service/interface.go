package service

import (
	"context"
	"go-server-exmaple/model"
)

type TodoService interface {
	GetTodo(ctx context.Context, id int) (*model.Todo, error)
	GetTodos(ctx context.Context) ([]model.Todo, error)
	CreateTodo(ctx context.Context, todo model.Todo) (int, error)
	UpdateTodo(ctx context.Context, id int) error
	DeleteTodo(ctx context.Context, id int) error
}
