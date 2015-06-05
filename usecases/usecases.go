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

func (interactor *TodoInteractor) Create(todo *domain.Todo) error {
	err := interactor.repo.Create(todo)
	if err != nil {
		return err
	}
	return nil
}

func (interactor *TodoInteractor) FindAll() ([]*domain.Todo, error) {
	todos, err := interactor.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (interactor *TodoInteractor) FindById(id string) (*domain.Todo, error) {
	todo, err := interactor.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (interactor *TodoInteractor) DeleteById(id string) error {
	err := interactor.repo.DeleteById(id)
	if err != nil {
		return err
	}
	return nil
}
