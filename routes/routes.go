package routes

import (
	"back/controller"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func Route() *mux.Router {
	route := mux.NewRouter()
	route.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Basic GET endpoint"))
		json.NewEncoder(w).Encode("Hello");
	}).Methods("GET");
	route.HandleFunc("/api/mail",controller.SendData).Methods("POST")
return route
}