package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"main/connections"
	"main/routes"
)

func main() {
	connections.Migrate()

	router := mux.NewRouter()
	routes.SetUserRoutes(router)
	connections.EnableCORS(router)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Server running on port 8080")
	log.Println(server.ListenAndServe())
}
