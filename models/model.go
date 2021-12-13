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
	CreatedAtt time.Time
	CreatedAt int `json:"created_at" gorm:"coloumn:created_at,autoCreateTime"`
    UpdatedAt time.Time `json:"updated_at" gorm:"coloumn:updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index"`
}

type GetTodoRequest struct{
	Id string
}

