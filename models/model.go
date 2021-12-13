package models

import (
	"time"

	"gorm.io/gorm"
)

type PostTodoRequest struct{
	Id        string `json:"id" gorm:"coloumn:id"`
	Todo string `json:"todo" gorm:"column:todo"`
	CreatedAt time.Time `json:"created_at" gorm:"coloumn:created_at"`
    UpdatedAt time.Time `gorm:"updated_at" gorm:"coloumn:updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index"`
}

type GetTodoRequest struct{
	Id string
}
