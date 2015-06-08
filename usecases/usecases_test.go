package usecases

import (
	"testing"

	"github.com/arnoldcano/teaxdeax/domain"
)

var (
	todoRepositoryStub *TodoRepositoryStub
	todosInteractor    *TodosInteractor
)

type TodoRepositoryStub struct{}

func (r *TodoRepositoryStub) Create(todo *domain.Todo) error {
	return nil
}

func (r *TodoRepositoryStub) FindAll() ([]*domain.Todo, error) {
	todo := domain.NewTodo("123", "test")
	return []*domain.Todo{todo}, nil
}

func (r *TodoRepositoryStub) FindById(id string) (*domain.Todo, error) {
	todo := domain.NewTodo("123", "test")
	return todo, nil
}

func (r *TodoRepositoryStub) Update(todo *domain.Todo) error {
	return nil
}

func (r *TodoRepositoryStub) DeleteById(id string) error {
	return nil
}

func init() {
	todosInteractor = NewTodosInteractor(todoRepositoryStub)
}

func TestNewTodosInteractor(t *testing.T) {
	if todosInteractor.repo != todoRepositoryStub {
		t.Errorf("Expected '%v', got '%v'", todoRepositoryStub, todosInteractor.repo)
	}
}

func TestTodosInteractorCreate(t *testing.T) {
	todo := &domain.Todo{}
	if err := todosInteractor.Create(todo); err != nil {
		t.Errorf("Expected '%v', got '%v'", nil, err)
	}
}

func TestTodosInteractorFindAll(t *testing.T) {
	todos, err := todosInteractor.FindAll()
	if err != nil {
		t.Errorf("Expected '%v', got '%v'", nil, err)
	}
	if len(todos) != 1 {
		t.Errorf("Expected '%v', got '%v'", 1, len(todos))
	}
}

func TestTodosInteractorFindById(t *testing.T) {
	todo, err := todosInteractor.FindById("123")
	if err != nil {
		t.Errorf("Expected '%v', got '%v'", nil, err)
	}
	if todo == nil {
		t.Errorf("Expected '%v', got '%v'", todo, nil)
	}
}

func TestTodosInteractorUpdate(t *testing.T) {
	todo := &domain.Todo{}
	err := todosInteractor.Update(todo)
	if err != nil {
		t.Errorf("Expected '%v', got '%v'", nil, err)
	}
}

func TestTodosInteractorDeleteById(t *testing.T) {
	err := todosInteractor.DeleteById("123")
	if err != nil {
		t.Errorf("Expected '%v', got '%v'", nil, err)
	}
}
