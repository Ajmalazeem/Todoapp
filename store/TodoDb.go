package store

import (
	"github.com/Ajmalazeem/models"
	"gorm.io/gorm"
)

type TodoStore interface {
	PostTodo(models.PostTodoRequest) error
	GetTodo(models.GetTodoRequest) (*models.PostTodoRequest, error)
}

type todoStore struct {
	db *gorm.DB
}

func (t *todoStore) PostTodo(req models.PostTodoRequest) error {
	return t.db.Table("todoer").Create(&req).Error
}

func (t *todoStore) GetTodo(req models.GetTodoRequest) (*models.PostTodoRequest, error) {
	result := models.PostTodoRequest{}
	err := t.db.Table("todoer").Where("id = ?", req.Id).Find(&result).Error

	if err != nil {
		return nil, err
	}

	return &result, nil

}

func NewTodoStore(db *gorm.DB) TodoStore {
	return &todoStore{
		db: db,
	}
}
