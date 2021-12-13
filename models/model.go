package models

import ("time"

 	//"gorm.io/gorm"
)

type PostTodoRequest struct{
	ID        string `json:"id" gorm:"coloumn:ID"`
	Todo string `json:"todo" gorm:"column:todo"`
	CreatedAt time.Time
    UpdatedAt time.Time
}

type GetTodoRequest struct{
	ID string
}