package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Ajmalazeem/store"
	"github.com/Ajmalazeem/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Ajmalazeem/Api"
)

func main(){
	config , err := util.LoadConfig(".")
	if err != nil{
		log.Fatal("cannot load config")
	}
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",config.DBHost,config.DBPort,config.DBUser,config.DBPassword,config.DBName)
	db , err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		panic(err)
	}
	todoStore := store.NewTodoStore(db)

	todo := api.NewTodoService(todoStore)
	log.Println("Listening on", "8000")
	http.ListenAndServe(":8000", api.MakeHandler(todo))
}