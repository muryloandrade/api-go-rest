package routes

import (
	"api-go-rest/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleResquests() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/personalidades", controllers.CriaPersonalidade).Methods("Post")
	r.HandleFunc("/personalidades/{id}", controllers.RetornaID).Methods("Get")
	r.HandleFunc("/personalidades/{id}", controllers.DeletaPersonalidade).Methods("Delete")

	log.Fatal(http.ListenAndServe(":8080", r))
}
