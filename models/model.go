package models

import (
	//"time"
	"github.com/Ajmalazeem/tim"

	"gorm.io/gorm"
)

type PostTodoRequest struct{
	Id        string `json:"id" gorm:"coloumn:id"`
	Todo string `json:"todo" gorm:"column:todo"`
	CreatedAt tim.Tim`json:"created_at" gorm:"coloumn:created_at"`
    UpdatedAt tim.Tim `json:"updated_at" gorm:"coloumn:updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index"`
}

type GetTodoRequest struct{
	Id string
}

