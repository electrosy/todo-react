package usecases_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/electrosy/todo-react/backend/entities"
	"github.com/electrosy/todo-react/backend/usecases"
)

var dummyTodos = []entities.Todo{
	{
		Title:       "Todo 1",
		Description: "description of todo 1",
		IsCompleted: true,
	},
	{
		Title:       "Todo 2",
		Description: "description of todo 2",
		IsCompleted: false,
	},
	{
		Title:       "Todo 3",
		Description: "description of todo 3",
		IsCompleted: true,
	},
}

type BadTodosRepo struct{}

func (BadTodosRepo) GetAllTodos() ([]entities.Todo, error) {
	return nil, fmt.Errorf("something went wrong")
}

type MockTodosRepo struct{}

func (MockTodosRepo) GetAllTodos() ([]entities.Todo, error) {
	return dummyTodos, nil
}

func TestGetTodos(t *testing.T) {
	//Test
	t.Run("Returns ErrInternal when TodosRepository returns error", func(t *testing.T) {

		repo := new(BadTodosRepo)

		todos, err := usecases.GetTodos(repo)

		if err != usecases.ErrInternal {
			t.Fatalf("Expected ErrInternal: Got: %v", err)
		}
		if todos != nil {
			t.Fatalf("Expected todos to be nil: Got: %v", todos)
		}
	})

	//Test
	t.Run("Returns todos from  TodosRepository", func(t *testing.T) {
		repo := new(MockTodosRepo)

		todos, err := usecases.GetTodos(repo)

		if !reflect.DeepEqual(todos, dummyTodos) {
			t.Fatalf("Expected todos to be dummyTodos: Got: %v", todos)
		}

		if err != nil {
			t.Fatalf("Expected error to be nill: Got: %v", err)
		}
	})
}
