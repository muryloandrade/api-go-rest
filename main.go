package main

import (
	"api-go-rest/models"
	"api-go-rest/routes"
	"fmt"
)

func main() {

	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Joao", Historia: "Joao era um menino muito legal"},
		{Id: 2, Nome: "Maria", Historia: "Maria era uma menina muito legal"},
	}

	fmt.Println("Iniciando servidor, redirecionando para home...")

	routes.HandleResquests()
}
