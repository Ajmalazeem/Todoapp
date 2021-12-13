package models

import (
 	"gorm.io/gorm"
)

type PostTodoRequest struct{
	gorm.Model
	ID        string `json:"ID" gorm:"coloumn:ID"`
	Todo string `json:"todo" gorm:"column:todo"`
}

type GetTodoRequest struct{
	ID string
}