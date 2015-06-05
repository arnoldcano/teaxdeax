package domain

type TodoRepository interface {
	Create(*Todo) error
	FindAll() ([]*Todo, error)
	FindById(string) (*Todo, error)
	Update(*Todo) error
	DeleteById(string) error
}

type Todo struct {
	Id   string `json:"id"`
	Note string `json:"note"`
}

func NewTodo(id, note string) *Todo {
	return &Todo{
		Id:   id,
		Note: note,
	}
}
