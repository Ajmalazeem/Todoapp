package models

type PostTodoRequest struct{
	Id string `json:"id" gorm:"column:id"`
	Todo string `json:"todo" gorm:"column:todo"`
}

type GetTodoRequest struct{
	Id string
}