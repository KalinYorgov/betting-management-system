package routes

import (
	"myapp/handlers"

	"github.com/gorilla/mux"
)

// InitializeRoutes sets up the routes for the application
func InitializeRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/entries", handlers.GetAllEntries).Methods("GET")
	router.HandleFunc("/entry", handlers.CreateEntry).Methods("POST")
	return router
}
