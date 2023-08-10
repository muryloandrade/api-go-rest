package main

import (
	"api-go-rest/database"
	"api-go-rest/routes"
	"fmt"
)

func main() {

	database.DbTwoConnect()

	fmt.Println("Iniciando servidor, redirecionando para home...")

	routes.HandleResquests()
}
