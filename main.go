package main

import (
	"fmt"

	"github.com/thiagocdn/alura-api-go-rest/routes"
)

func main() {
	fmt.Println("Iniciando servidor REST com Go...")
	routes.HandleRequest()
}
