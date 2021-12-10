package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	//"ggrc/models"
	"ggrc/store"
	"ggrc/Api"
)

func main(){
	db , err := gorm.Open(postgres.Open("host=localhost user=postgres password=12345 dbname=todoer port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil{
		panic(err)
	}
	todoStore := store.NewTodoStore(db)

	todo := api.NewTodoService(todoStore)
	log.Println("Listening on", "8000")
	http.ListenAndServe(":8000", api.MakeHandler(todo))
}