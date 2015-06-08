package domain

import "testing"

func TestNewTodo(t *testing.T) {
	id := "123"
	note := "new note"
	todo := NewTodo(id, note)
	if todo.Id != id {
		t.Errorf("Expected '%v', got '%v'", id, todo.Id)
	}
	if todo.Note != note {
		t.Errorf("Expected '%v', got '%v'", note, todo.Note)
	}
}
