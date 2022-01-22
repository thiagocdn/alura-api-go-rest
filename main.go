package main

import (
	"fmt"

	"github.com/thiagocdn/alura-api-go-rest/models"
	"github.com/thiagocdn/alura-api-go-rest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Nome: "Thiago", Historia: "Historia muito boa"},
		{Nome: "Thiago 2", Historia: "Historia muito boa 2"},
	}

	fmt.Println("Iniciando servidor REST com Go...")
	routes.HandleRequest()
}
