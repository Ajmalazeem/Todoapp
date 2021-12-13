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
	CreatedAt time.Time`json:"created_at" gorm:"coloumn:created_at"`
	CreatedAtt string
    UpdatedAt time.Time `json:"updated_at" gorm:"coloumn:updated_at"`
	UpdatedAtt string
    DeletedAt gorm.DeletedAt `gorm:"index"`
}

type GetTodoRequest struct{
	Id string
}

