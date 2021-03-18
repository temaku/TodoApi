package todo

import "github.com/temaku/TodoApi/model"

type TodoService interface {
	Todos() ([]model.Todo, []error)
	Todo(id uint) (*model.Todo, []error)
	UpdateTodo(todo *model.Todo) (*model.Todo, []error)
	DeleteTodo(id uint) (*model.Todo, []error)
	StoreTodo(todo *model.Todo) (*model.Todo, []error)
	TodoExists(name string) bool

	
}
