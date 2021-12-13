package models

import "gorm.io/gorm"

type PostTodoRequest struct{
	gorm.Model
	Id string `json:"id" gorm:"column:id"`
	Todo string `json:"todo" gorm:"column:todo"`
}

type GetTodoRequest struct{
	Id string
}