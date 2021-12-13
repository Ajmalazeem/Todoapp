package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"github.com/Ajmalazeem/store"
	"github.com/Ajmalazeem/util"
	
	"github.com/Ajmalazeem/Api"
)

func main(){
	db , err := gorm.Open(postgres.Open(), &gorm.Config{})
	if err != nil{
		panic(err)
	}
	todoStore := store.NewTodoStore(db)

	todo := api.NewTodoService(todoStore)
	log.Println("Listening on", "8000")
	http.ListenAndServe(":8000", api.MakeHandler(todo))
}