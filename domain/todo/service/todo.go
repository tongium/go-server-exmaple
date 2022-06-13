package service

import (
	"context"
	"go-server-exmaple/domain/todo/repository"
	"go-server-exmaple/model"
)

type todoService struct {
	todoRepository repository.TodoRepository
}

// CreateTodo implements TodoService
func (s *todoService) CreateTodo(ctx context.Context, todo model.Todo) (int, error) {
	if todo.Task == "" {
		return 0, ErrInvalidTask
	}

	return s.todoRepository.CreateTodo(ctx, todo)
}

// DeleteTodo implements TodoService
func (s *todoService) DeleteTodo(ctx context.Context, id int) error {
	panic("unimplemented")
}

// GetTodo implements TodoService
func (s *todoService) GetTodo(ctx context.Context, id int) (*model.Todo, error) {
	return s.todoRepository.GetTodo(ctx, id)
}

// GetTodos implements TodoService
func (s *todoService) GetTodos(ctx context.Context) ([]model.Todo, error) {
	return s.todoRepository.GetTodos(ctx)
}

// UpdateTodo implements TodoService
func (s *todoService) UpdateTodo(ctx context.Context, id int) error {
	panic("unimplemented")
}

func NewTodoService(todoRepository repository.TodoRepository) TodoService {
	return &todoService{todoRepository: todoRepository}
}
