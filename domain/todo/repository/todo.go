package repository

import (
	"context"
	"go-server-exmaple/model"

	"gorm.io/gorm"
)

type todoRepository struct {
	db *gorm.DB
}

// CreateTodo implements TodoRepository
func (r *todoRepository) CreateTodo(ctx context.Context, todo model.Todo) (int, error) {
	result := r.db.WithContext(ctx).Create(&todo)
	return todo.ID, result.Error
}

// DeleteTodo implements TodoRepository
func (r *todoRepository) DeleteTodo(ctx context.Context, id int) error {
	panic("unimplemented")
}

// GetTodo implements TodoRepository
func (r *todoRepository) GetTodo(ctx context.Context, id int) (*model.Todo, error) {
	var todo model.Todo
	result := r.db.WithContext(ctx).Where("id = ?", id).First(&todo)
	return &todo, result.Error
}

// GetTodos implements TodoRepository
func (r *todoRepository) GetTodos(ctx context.Context) ([]model.Todo, error) {
	var todos []model.Todo
	result := r.db.WithContext(ctx).Find(&todos)
	return todos, result.Error
}

// UpdateTodo implements TodoRepository
func (r *todoRepository) UpdateTodo(ctx context.Context, todo model.Todo) error {
	panic("unimplemented")
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{db: db}
}
