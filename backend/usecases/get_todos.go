package usecases

import "github.com/electrosy/todo-react/backend/entities"

func GetTodos(repo TodosRepository) ([]entities.Todo, error) {
	return nil, ErrInternal
}
