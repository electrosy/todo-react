package usecases_test

import (
	"testing"
	"fmt"

	"github.com/electrosy/todo-react/backend/entities"
	"github.com/electrosy/todo-react/backend/usecases"
)

type MockTodosRepo struct {}

func (MockTodosRepo) GetAllTodos() ([]entities.Todo, error) {
	return nil, fmt.Errorf("something went wrong")	
}

func TestGetTodos(t *testing.T) {
	repo := new(MockTodosRepo)

	todos, err := usecases.GetTodos(repo)

	if err != usecases.ErrInternal {
		t.Fatalf("Expected ErrInternal: Got: %v", err)
	}
	if todos != nil {
		t.Fatalf("Expected todos to be nil: Got: %v", todos)
	}
}