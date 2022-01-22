package main

import (
	"fmt"

	"github.com/thiagocdn/alura-api-go-rest/database"
	"github.com/thiagocdn/alura-api-go-rest/models"
	"github.com/thiagocdn/alura-api-go-rest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Thiago", Historia: "Historia muito boa"},
		{Id: 2, Nome: "Thiago 2", Historia: "Historia muito boa 2"},
	}

	database.ConectaComBancoDeDados()

	fmt.Println("Iniciando servidor REST com Go...")
	routes.HandleRequest()
}
