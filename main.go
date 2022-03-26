package main

import (
	"api-go/db"
	"api-go/routes"
)

func main() {
	database.ConecaoCOmBancoDeDados()
	routes.HandleRequests()
}
