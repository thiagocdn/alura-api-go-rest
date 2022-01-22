package routes

import (
	"log"
	"net/http"

	"github.com/thiagocdn/alura-api-go-rest/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalidades", controllers.TodasPersonalidades)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
