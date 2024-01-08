package main

import (
	"myapp/routes"
	"net/http"
)

func main() {
	router := routes.InitializeRoutes()
	http.ListenAndServe(":8080", router)
}
