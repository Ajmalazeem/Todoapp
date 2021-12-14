package api

import (
	"github.com/Ajmalazeem/models"
	"github.com/Ajmalazeem/store"
)

type Todo interface {
	PostTodo(models.PostTodoRequest)
	GetTodo(models.GetTodoRequest)(*models.PostTodoRequest, error)
	// PutTodo(models.PutTodoRequest)
	DeleteTodo(models.DeleteTodoRequest)
}

type todo struct{ 
	todoStore store.TodoStore
}

func (t *todo) PostTodo(req models.PostTodoRequest){
	t.todoStore.PostTodo(req)
}

func (t *todo) GetTodo(req models.GetTodoRequest)(*models.PostTodoRequest,error){
	return t.todoStore.GetTodo(req)
}

// func (t *todo) PutTodo(req models.PutTodoRequest){
	// t.todoStore.PutTodo(req)
// }

func (t *todo) DeleteTodo(req models.DeleteTodoRequest){
	t.todoStore.DeleteTodo(req)
}

func NewTodoService(todoStore store.TodoStore) Todo{
	return &todo{
		todoStore: todoStore,
	}
}






