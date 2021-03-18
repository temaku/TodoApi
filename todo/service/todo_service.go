package services

import (
	todo "github.com/temaku/TodoApi/todo"
	"github.com/temaku/TodoApi/model"

)

type TodoService struct {
	todoRepo todo.TodoRepository
}

func NewTodoService(todoRep todo.TodoRepository) *TodoService {
	return &TodoService{todoRepo: todoRep}
}

func (cs *TodoService) Todos() ([]model.Todo, []error) {

	todos, errs := cs.todoRepo.Todos()
	if len(errs) > 0 {
		return nil, errs
	}
	return todos, errs

}


func (cs *TodoService) Todo(id uint) (*model.Todo, []error) {
	cl, errs := cs.todoRepo.Todo(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return cl, errs

}

func (cs *TodoService) UpdateTodo(todo *model.Todo) (*model.Todo, []error) {
	cl, errs := cs.todoRepo.UpdateTodo(todo)
	if len(errs) > 0 {
		return nil, errs
	}
	return cl, errs

}

func (cs *TodoService) DeleteTodo(id uint) (*model.Todo, []error) {

	cl, errs := cs.todoRepo.DeleteTodo(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return cl, errs
}

func (cs *TodoService) StoreTodo(todo *model.Todo) (*model.Todo, []error) {

	cl, errs := cs.todoRepo.StoreTodo(todo)
	if len(errs) > 0 {
		return nil, errs
	}
	return cl, errs
}
func (cs *TodoService) TodoExists(todoname string) bool {
	exists := cs.todoRepo.TodoExists(todoname)
	return exists
}
