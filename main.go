package main

import (
	"api-go-rest/routes"
	"fmt"
)

func main() {
	fmt.Println("Iniciando servidor, redirecionando para home...")
	routes.HandleResquests()
}
