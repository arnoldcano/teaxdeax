package interfaces

import (
	"testing"

	"github.com/arnoldcano/teaxdeax/domain"
)

var (
	databaseStub    *DatabaseStub
	rowsStub        *RowsStub
	todosRepository *TodosRepository
)

type DatabaseStub struct{}

func (d *DatabaseStub) Execute(query string) error {
	return nil
}

func (d *DatabaseStub) Query(query string) (Rows, error) {
	return rowsStub, nil
}

type RowsStub struct{}

func (r *RowsStub) Scan(dest ...interface{}) error {
	return nil
}

func (r *RowsStub) Next() bool {
	return false
}

func (r *RowsStub) Close() error {
	return nil
}

func (r *RowsStub) Err() error {
	return nil
}

func init() {
	todosRepository = NewTodosRepository(databaseStub)
}

func TestNewTodosRepository(t *testing.T) {
	if todosRepository.db != databaseStub {
		t.Errorf("Expected '%v', got '%v'", databaseStub, todosRepository.db)
	}
}

func TestTodosRepositoryCreate(t *testing.T) {
	todo := &domain.Todo{}
	if err := todosRepository.Create(todo); err != nil {
		t.Errorf("Expected '%v', got '%v'", nil, err)
	}
}

func TestTodosRepositoryFindAll(t *testing.T) {
	todos, err := todosRepository.FindAll()
	if err != nil {
		t.Errorf("Expected '%v', got '%v'", nil, err)
	}
	if len(todos) != 0 {
		t.Errorf("Expected '%v', got '%v'", 0, len(todos))
	}
}

func TestTodosRepositoryFindById(t *testing.T) {
	todo, err := todosRepository.FindById("123")
	if err != nil {
		t.Errorf("Expected '%v', got '%v'", nil, err)
	}
	if todo == nil {
		t.Errorf("Expected '%v', got '%v'", todo, nil)
	}
}

func TestTodosRepositoryUpdate(t *testing.T) {
	todo := &domain.Todo{}
	err := todosRepository.Update(todo)
	if err != nil {
		t.Errorf("Expected '%v', got '%v'", nil, err)
	}
}

func TestTodosRepositoryDeleteById(t *testing.T) {
	err := todosRepository.DeleteById("123")
	if err != nil {
		t.Errorf("Expected '%v', got '%v'", nil, err)
	}
}
