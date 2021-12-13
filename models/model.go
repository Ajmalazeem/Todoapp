package models

import (
	"time"

	//"gorm.io/gorm"
)

type PostTodoRequest struct{
	//gorm.Model
	Id        string `json:"id" gorm:"coloumn:id"`
	Todo string `json:"todo" gorm:"column:todo"`
	CreatedAt *time.Time
    UpdatedAt *time.Time
    //DeletedAt DeletedAt `gorm:"index"`
}

type GetTodoRequest struct{
	Id string
}