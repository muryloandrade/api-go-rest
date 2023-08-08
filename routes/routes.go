package routes

import (
	"api-go-rest/controllers"
	"api-go-rest/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleResquests() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/personalidades", controllers.CriaPersonalidade).Methods("Post")
	r.HandleFunc("/personalidades/{id}", controllers.RetornaID).Methods("Get")
	r.HandleFunc("/personalidades/{id}", controllers.DeletaPersonalidade).Methods("Delete")
	r.HandleFunc("/personalidades/{id}", controllers.EditPersonalidade).Methods("Put")
	log.Fatal(http.ListenAndServe(":8080", r))
}
