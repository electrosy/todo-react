package usecases

import "github.com/electrosy/todo-react/backend/entities"

type TodosRepository interface {
	GetAllTodos() ([]entities.Todo, error)
}
