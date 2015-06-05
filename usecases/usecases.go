package usecases

import "github.com/arnoldcano/teaxdeax/domain"

type TodoInteractor struct {
	repo domain.TodoRepository
}

func NewTodoInteractor(repo domain.TodoRepository) *TodoInteractor {
	return &TodoInteractor{
		repo: repo,
	}
}

func (i *TodoInteractor) Create(todo *domain.Todo) error {
	err := i.repo.Create(todo)
	if err != nil {
		return err
	}
	return nil
}

func (i *TodoInteractor) FindAll() ([]*domain.Todo, error) {
	todos, err := i.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (i *TodoInteractor) FindById(id string) (*domain.Todo, error) {
	todo, err := i.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (i *TodoInteractor) Update(todo *domain.Todo) error {
	err := i.repo.Update(todo)
	if err != nil {
		return err
	}
	return nil
}

func (i *TodoInteractor) DeleteById(id string) error {
	err := i.repo.DeleteById(id)
	if err != nil {
		return err
	}
	return nil
}
