package interfaces

import (
	"fmt"

	"github.com/arnoldcano/teaxdeax/domain"
)

type Database interface {
	Execute(query string) error
	Query(query string) (Rows, error)
}

type Rows interface {
	Scan(dest ...interface{}) error
	Next() bool
	Close() error
}

type TodosRepository struct {
	db Database
}

func NewTodosRepository(db Database) *TodosRepository {
	return &TodosRepository{
		db: db,
	}
}

func (r *TodosRepository) Create(todo *domain.Todo) error {
	query := fmt.Sprintf(
		"INSERT INTO todos (id, note) VALUES ('%v', '%v')",
		todo.Id,
		todo.Note,
	)
	if err := r.db.Execute(query); err != nil {
		return err
	}
	return nil
}

func (r *TodosRepository) FindAll() ([]*domain.Todo, error) {
	var (
		todos []*domain.Todo
		id    string
		note  string
	)
	rows, err := r.db.Query("SELECT id, note FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &note)
		if err != nil {
			return nil, err
		}
		todo := domain.NewTodo(id, note)
		todos = append(todos, todo)
	}
	return todos, nil
}

func (r *TodosRepository) FindById(id string) (*domain.Todo, error) {
	var note string
	query := fmt.Sprintf("SELECT note FROM todos WHERE id = '%v'", id)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	rows.Next()
	if err = rows.Scan(&note); err != nil {
		return nil, err
	}
	todo := domain.NewTodo(id, note)
	return todo, nil
}

func (r *TodosRepository) Update(todo *domain.Todo) error {
	query := fmt.Sprintf("UPDATE todos SET note='%v' WHERE id='%v'", todo.Note, todo.Id)
	if err := r.db.Execute(query); err != nil {
		return err
	}
	return nil
}

func (r *TodosRepository) DeleteById(id string) error {
	query := fmt.Sprintf("DELETE FROM todos WHERE id='%v'", id)
	if err := r.db.Execute(query); err != nil {
		return err
	}
	return nil
}
