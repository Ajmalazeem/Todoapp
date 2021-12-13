package store

import (
	"time"

	"github.com/Ajmalazeem/models"
	"gorm.io/gorm"
)

type TodoStore interface{
	PostTodo(models.PostTodoRequest)error
	GetTodo(models.GetTodoRequest)(*models.PostTodoRequest, error)
}

type todoStore struct{
	db *gorm.DB
}

func(t *todoStore) PostTodo(req models.PostTodoRequest)error{
	req.CreatedAt = time.Now()
	req.UpdatedAt = time.Now()

	return t.db.Table("todoer").Create(&req).Error
	
}
// func tim (t time.Time) string{
// 	x := t.String()
// 	return x
// }

func(t *todoStore) GetTodo(req models.GetTodoRequest)(*models.PostTodoRequest,error){
	var result models.PostTodoRequest
	t.db.Table("todoer").Where("id = ?",req.Id).Find(&result)
	
	return &result, nil

}

func NewTodoStore(db *gorm.DB) TodoStore {
	return &todoStore{
		db: db,
	}
}
