package interfaces

import (
	"fmt"

	"github.com/arnoldcano/teaxdeax/domain"
)

type Database interface {
	Execute(statement string) error
	Query(statement string) (Rows, error)
}

type Rows interface {
	Scan(dest ...interface{}) error
	Next() bool
	Close() error
}

type TodoRepository struct {
	db Database
}

func NewTodoRepository(db Database) *TodoRepository {
	return &TodoRepository{
		db: db,
	}
}

func (r *TodoRepository) Create(todo *domain.Todo) error {
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

func (r *TodoRepository) FindAll() ([]*domain.Todo, error) {
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

func (r *TodoRepository) FindById(id string) (*domain.Todo, error) {
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

func (r *TodoRepository) Update(todo *domain.Todo) error {
	query := fmt.Sprintf("UPDATE todos SET note='%v' WHERE id='%v'", todo.Note, todo.Id)
	if err := r.db.Execute(query); err != nil {
		return err
	}
	return nil
}

func (r *TodoRepository) DeleteById(id string) error {
	query := fmt.Sprintf("DELETE FROM todos WHERE id='%v'", id)
	if err := r.db.Execute(query); err != nil {
		return err
	}
	return nil
}
