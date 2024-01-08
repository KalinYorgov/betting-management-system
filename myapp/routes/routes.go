package routes

import (
	"myapp/handlers"

	"github.com/gorilla/mux"
)

func InitializeRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/entry", handlers.CreateEntry).Methods("POST")
	// Define other routes
	return router
}
