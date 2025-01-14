package routes

import (
	"back/controller"
	

	"github.com/gorilla/mux"
)

func Route() *mux.Router {
	route := mux.NewRouter()
	
	route.HandleFunc("/api/mail",controller.SendData).Methods("POST")
return route
}