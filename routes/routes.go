package routes

import (
	"backend_tarot/controllers"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Route struct {
	Path    string
	Handler http.HandlerFunc
	Method  string
}

// GetRoutes returns a slice of routes, accepting the database connection and passing it to handlers
func GetRoutes(db *gorm.DB) []Route {
	return []Route{
		// POST request to generate a tarot card
		{Path: "/api/generateTarotCard", Handler: controllers.GenerateTarotCardHandler(db), Method: "POST"},

		// GET request to get all available tarot cards
		{Path: "/api/getTarotCardsList", Handler: controllers.GetAllTarotCardsHandler(db), Method: "GET"},
	}
}

// LoadRoutes sets up the routes in the router with the provided database connection
func LoadRoutes(router *mux.Router, db *gorm.DB) {
	routes := GetRoutes(db)

	for _, route := range routes {
		router.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	}
}
