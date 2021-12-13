package models

import (
	//"time"
	//"github.com/Ajmalazeem/tim"

	"time"

	"gorm.io/gorm"
)

type PostTodoRequest struct{
	Id        string `json:"id" gorm:"coloumn:id"`
	Todo string `json:"todo" gorm:"column:todo"`
	CreatedAt time.Time`json:"created_at,omitempty" gorm:"coloumn:created_at"`
    UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"coloumn:updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index"`
}

type GetTodoRequest struct{
	Id string
}

