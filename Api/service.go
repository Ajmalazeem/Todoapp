package api

import (
	"ggrc/models"
	"ggrc/store"
)

type Todo interface {
	PostTodo(models.PostTodoRequest)
	GetTodo(models.GetTodoRequest)(*models.PostTodoRequest, error)
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

func NewTodoService(todoStore store.TodoStore) Todo{
	return &todo{
		todoStore: todoStore,
	}
}






