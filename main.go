package main

import (
	"api-go-rest/models"
	"api-go-rest/routes"
	"fmt"
)

func main() {

	models.Personalidades = []models.Personalidade{
		{Nome: "Joao", Historia: "Joao era um menino muito legal"},
		{Nome: "Maria", Historia: "Maria era uma menina muito legal"},
	}

	fmt.Println("Iniciando servidor, redirecionando para home...")

	routes.HandleResquests()
}
