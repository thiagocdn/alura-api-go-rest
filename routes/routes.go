package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thiagocdn/alura-api-go-rest/controllers"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades)
	log.Fatal(http.ListenAndServe(":3000", r))
}
