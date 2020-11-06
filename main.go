package main

import (
	"github.com/benacook/getGround-technical-task/controller"
	"github.com/benacook/getGround-technical-task/model/database"
	"log"
	"net/http"
)

func main() {
	if err := database.Init(); err != nil{
		log.Fatal(err)
	}
	defer database.Db.Close()

	controller.RegisterHandlers()
	log.Println("running on port 8080")
	if err :=  http.ListenAndServe(":8080", nil); err != nil{
		log.Fatal(err)
	}
}
