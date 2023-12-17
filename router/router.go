package router

import (
	"crud-app/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()
	//Router Handles Enpoints
	router.HandleFunc("/getAllUser", controller.GetAllUser).Methods("GET")
	//router.HandleFunc("/getUser/{id}", getUserbyId).Methods("GET")
	router.HandleFunc("/createUser", controller.CreateUser).Methods("POST")
	router.HandleFunc("/updateUser/{id}", controller.UpdateUser).Methods("PUT")
	router.HandleFunc("/deleteUser/{id}", controller.DeleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))

	return router

}
