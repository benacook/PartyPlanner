package main

import (
	"github.com/benacook/GetGround-Assignment/controller"
	"github.com/benacook/GetGround-Assignment/model/database"
	"log"
	"net/http"
)

func main() {
	if err := database.Init(); err != nil{
		log.Fatal(err)
	}
	defer database.Db.Close()

	controller.RegisterHandlers()
	log.Println("running on port 8000")
	if err :=  http.ListenAndServeTLS(":8000",
		"server.crt", "server.key", nil); err != nil{
		database.Db.Close()
		log.Fatal(err)
	}
}
