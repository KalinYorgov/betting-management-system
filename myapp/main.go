package main

import (
	"log"
	"myapp/routes"
	"net/http"
)

func main() {
	router := routes.InitializeRoutes()

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
