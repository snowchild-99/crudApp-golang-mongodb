package main

import (
	"crud-app/controller"
	"crud-app/router"
	"log"
	"net/http"
)

func main() {
	log.Println("Inside Main")

	//it will listen to the port number of router
	controller.Init()
	r := router.Router()

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Server Connection established")

}
