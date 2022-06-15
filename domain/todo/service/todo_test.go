package service_test

import (
	"context"
	"go-server-exmaple/domain/todo/service"
	"go-server-exmaple/mocks"
	"go-server-exmaple/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTodoWhenTaskIsValid(t *testing.T) {
	repo := &mocks.TodoRepository{}
	repo.On("CreateTodo", mock.Anything, mock.Anything).Return(1, nil)

	svc := service.NewTodoService(repo)
	_, err := svc.CreateTodo(context.Background(), model.Todo{Task: "todo"})
	assert.NoError(t, err)
}

func BenchmarkCreateTodo(b *testing.B) {
	repo := &mocks.TodoRepository{}
	repo.On("CreateTodo", mock.Anything, mock.Anything).Return(1, nil)

	svc := service.NewTodoService(repo)
	for i := 0; i < b.N; i++ {
		svc.CreateTodo(context.Background(), model.Todo{Task: "todo"})
	}
}
