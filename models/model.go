package models

import (
	//"time"
	//"github.com/Ajmalazeem/tim"

	"time"
)

type PostTodoRequest struct {
	Id        string     `json:"id" gorm:"coloumn:id"`
	Todo      string     `json:"todo" gorm:"column:todo"`
	CreatedAt time.Time  `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
}

type GetTodoRequest struct {
	Id string
}

// type PutTodoRequest struct{
	// Id string
	// Todo string
// }

type DeleteTodoRequest struct{
	Id string
}
