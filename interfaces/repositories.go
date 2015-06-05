package interfaces

import (
	"fmt"
	"time"

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

func (repo *TodoRepository) Create(todo *domain.Todo) error {
	query := fmt.Sprintf(
		"INSERT INTO todos (id, note, created_at, updated_at) VALUES ('%v', '%v', '%v', '%v')",
		todo.Id,
		todo.Note,
		todo.CreatedAt,
		todo.UpdatedAt,
	)
	if err := repo.db.Execute(query); err != nil {
		return err
	}
	return nil
}

func (repo *TodoRepository) FindAll() ([]*domain.Todo, error) {
	var todos []*domain.Todo
	var id string
	var note string
	var createdAt time.Time
	var updatedAt time.Time
	rows, err := repo.db.Query("SELECT id, note, created_at, updated_at FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &note, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}
		todo := domain.NewTodo(id, note, createdAt, updatedAt)
		todos = append(todos, todo)
	}
	return todos, nil
}

func (repo *TodoRepository) FindById(id string) (*domain.Todo, error) {
	var note string
	var createdAt time.Time
	var updatedAt time.Time
	query := fmt.Sprintf(
		"SELECT note, created_at, updated_at FROM todos WHERE id = '%v'", id,
	)
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	rows.Next()
	if err = rows.Scan(&note, &createdAt, &updatedAt); err != nil {
		return nil, err
	}
	todo := domain.NewTodo(id, note, createdAt, updatedAt)
	return todo, nil
}

func (repo *TodoRepository) DeleteById(id string) error {
	query := fmt.Sprintf("DELETE FROM todos WHERE id='%v'", id)
	if err := repo.db.Execute(query); err != nil {
		return err
	}
	return nil
}
