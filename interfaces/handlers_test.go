package interfaces

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/arnoldcano/teaxdeax/domain"
)

var (
	todosInteractorStub *TodosInteractorStub
	todosHandler        *TodosHandler
	todo                *domain.Todo
	req                 *http.Request
)

type TodosInteractorStub struct{}

func (s *TodosInteractorStub) Create(todo *domain.Todo) error {
	return nil
}

func (s *TodosInteractorStub) FindAll() ([]*domain.Todo, error) {
	return []*domain.Todo{todo}, nil
}

func (s *TodosInteractorStub) FindById(id string) (*domain.Todo, error) {
	return todo, nil
}

func (s *TodosInteractorStub) Update(todo *domain.Todo) error {
	return nil
}

func (s *TodosInteractorStub) DeleteById(id string) error {
	return nil
}

func init() {
	todosHandler = NewTodosHandler(todosInteractorStub)
	todo = domain.NewTodo("123", "test")
	req = &http.Request{}
}

func TestNewTodosHandler(t *testing.T) {
	if todosHandler.interactor != todosInteractorStub {
		t.Errorf("Expected '%v', got '%v'", todosInteractorStub, todosHandler.interactor)
	}
}

func TestTodosHandlerCreate(t *testing.T) {
	res := httptest.NewRecorder()
	todosHandler.Create(res, req)
	if res.Code != http.StatusOK {
		t.Errorf("Expected '%v', got '%v'", http.StatusOK, res.Code)
	}
}

func TestTodosHandlerIndex(t *testing.T) {
	res := httptest.NewRecorder()
	todosHandler.Index(res, req)
	if res.Code != http.StatusOK {
		t.Errorf("Expected '%v', got '%v'", http.StatusOK, res.Code)
	}
	header := "application/json"
	if !strings.Contains(res.HeaderMap["Content-Type"][0], header) {
		t.Errorf("Expected '%v', got '%v'", header, res.HeaderMap["Content-Type"])
	}
	json := "[{\"id\":\"123\",\"note\":\"test\"}]"
	if res.Body.String() != json {
		t.Errorf("Expected '%v', got '%v'", json, res.Body)
	}
}

func TestTodosHandlerShow(t *testing.T) {
	res := httptest.NewRecorder()
	todosHandler.Show(res, req)
	if res.Code != http.StatusOK {
		t.Errorf("Expected '%v', got '%v'", http.StatusOK, res.Code)
	}
	header := "application/json"
	if !strings.Contains(res.HeaderMap["Content-Type"][0], header) {
		t.Errorf("Expected '%v', got '%v'", header, res.HeaderMap["Content-Type"])
	}
	json := "{\"id\":\"123\",\"note\":\"test\"}"
	if res.Body.String() != json {
		t.Errorf("Expected '%v', got '%v'", json, res.Body)
	}
}

func TestTodosHandlerUpdate(t *testing.T) {
	res := httptest.NewRecorder()
	todosHandler.Update(res, req)
	if res.Code != http.StatusOK {
		t.Errorf("Expected '%v', got '%v'", http.StatusOK, res.Code)
	}
}

func TestTodosHandlerDelete(t *testing.T) {
	res := httptest.NewRecorder()
	todosHandler.Destroy(res, req)
	if res.Code != http.StatusOK {
		t.Errorf("Expected '%v', got '%v'", http.StatusOK, res.Code)
	}
}
