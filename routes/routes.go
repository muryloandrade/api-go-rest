package routes

import (
	"api-go-rest/controllers"
	"api-go-rest/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleResquests() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/banco", controllers.TodosUsuarios).Methods("Get")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
