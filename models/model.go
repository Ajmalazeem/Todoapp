package models

import (
 	"gorm.io/gorm"
)

type PostTodoRequest struct{
	gorm.Model
	ID        string `json:"id" gorm:"coloumn:id"`
	Todo string `json:"todo" gorm:"column:todo"`
}

type GetTodoRequest struct{
	ID string
}