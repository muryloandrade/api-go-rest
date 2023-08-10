package controllers

import (
	"api-go-rest/database"
	"api-go-rest/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bem vindo a rota Home")
}

func TodosUsuarios(w http.ResponseWriter, r *http.Request) {
	var per []models.User
	database.DBONE.Query("SELECT * FROM personalidades", &per)
	json.NewEncoder(w).Encode(per)
}
