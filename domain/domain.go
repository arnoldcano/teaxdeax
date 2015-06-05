package domain

import "time"

type TodoRepository interface {
	Create(*Todo) error
	FindAll() ([]*Todo, error)
	FindById(string) (*Todo, error)
	DeleteById(string) error
}

type Todo struct {
	Id        string    `json:"id"`
	Note      string    `json:"note"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewTodo(id, note string, createdAt, updatedAt time.Time) *Todo {
	return &Todo{
		Id:        id,
		Note:      note,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
