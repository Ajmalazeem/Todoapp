package api

import (
	"github.com/Ajmalazeem/models"
	"github.com/Ajmalazeem/store"
)

type Todo interface {
	PostTodo(models.PostTodoRequest) error
	GetTodo(models.GetTodoRequest)(*models.PostTodoRequest, error)
	PutTodo(models.PutTodoRequest) error
	DeleteTodo(models.DeleteTodoRequest) error
}

type todo struct{ 
	todoStore store.TodoStore
}

func (t *todo) PostTodo(req models.PostTodoRequest)error{
	t.todoStore.PostTodo(req)
	return nil
}

func (t *todo) GetTodo(req models.GetTodoRequest)(*models.PostTodoRequest,error){
	return t.todoStore.GetTodo(req)
}

func (t *todo) PutTodo(req models.PutTodoRequest)error{
	t.todoStore.PutTodo(req)
	return nil
}

func (t *todo) DeleteTodo(req models.DeleteTodoRequest)error{
	t.todoStore.DeleteTodo(req)
	return nil
}

func NewTodoService(todoStore store.TodoStore) Todo{
	return &todo{
		todoStore: todoStore,
	}
}






