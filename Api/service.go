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
	 if err :=t.todoStore.PostTodo(req); err!= nil{
		return err
	}
	return nil
}

func (t *todo) GetTodo(req models.GetTodoRequest)(*models.PostTodoRequest,error){
	d , err:= t.todoStore.GetTodo(req)
	if err != nil{
		return nil, err
	}
	return d, nil
}

func (t *todo) PutTodo(req models.PutTodoRequest)error{
	if err:=t.todoStore.PutTodo(req); err!= nil{
		return err
	}
	return nil
}

func (t *todo) DeleteTodo(req models.DeleteTodoRequest)error{
	if err:= t.todoStore.DeleteTodo(req); err!= nil{
		return err
	}
	return nil
}

func NewTodoService(todoStore store.TodoStore) Todo{
	return &todo{
		todoStore: todoStore,
	}
}






