package routes

import (
	"backend_tarot/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Path    string
	Handler http.HandlerFunc
	Method  string
}

func GetRoutes() []Route {
	return []Route{
		// POST request to generate a tarot card
		{Path: "/api/generateTarotCard", Handler: controllers.GenerateTarotCardHandler, Method: "POST"},

		// GET request to get all available tarot cards
		{Path: "/api/getTarotCardsList", Handler: controllers.GetAllTarotCardsHandler, Method: "GET"},
	}
}

func LoadRoutes(router *mux.Router) {
	routes := GetRoutes()

	for _, route := range routes {
		router.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	}
}
