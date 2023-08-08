package controllers

import (
	"api-go-rest/database"
	"api-go-rest/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bem vindo a rota Home")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	var per []models.Personalidade
	database.DB.Find(&per)
	json.NewEncoder(w).Encode(per)
}

func RetornaID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)
	json.NewEncoder(w).Encode(personalidade)
}

func CriaPersonalidade(w http.ResponseWriter, r *http.Request) {
	var newPersonalidade models.Personalidade
	json.NewDecoder(r.Body).Decode(&newPersonalidade)
	database.DB.Create(&newPersonalidade)
	json.NewEncoder(w).Encode(newPersonalidade)
}

func DeletaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidadeDelete models.Personalidade
	database.DB.Delete(&personalidadeDelete, id)
	json.NewEncoder(w).Encode(personalidadeDelete)
}
