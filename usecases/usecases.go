package usecases

import "github.com/arnoldcano/teaxdeax/domain"

type TodosInteractor struct {
	repo domain.TodosRepository
}

func NewTodosInteractor(repo domain.TodosRepository) *TodosInteractor {
	return &TodosInteractor{
		repo: repo,
	}
}

func (i *TodosInteractor) Create(todo *domain.Todo) error {
	if err := i.repo.Create(todo); err != nil {
		return err
	}
	return nil
}

func (i *TodosInteractor) FindAll() ([]*domain.Todo, error) {
	todos, err := i.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (i *TodosInteractor) FindById(id string) (*domain.Todo, error) {
	todo, err := i.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (i *TodosInteractor) Update(todo *domain.Todo) error {
	if _, err := i.repo.FindById(todo.Id); err != nil {
		return err
	}
	if err := i.repo.Update(todo); err != nil {
		return err
	}
	return nil
}

func (i *TodosInteractor) DeleteById(id string) error {
	if _, err := i.repo.FindById(id); err != nil {
		return err
	}
	if err := i.repo.DeleteById(id); err != nil {
		return err
	}
	return nil
}
