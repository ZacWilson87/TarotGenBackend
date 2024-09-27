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
		// Tarot Card Generation Routes
		{Path: "/api/generateTarotCard", Handler: controllers.GenerateTarotCardHandler(db), Method: "POST"},
		{Path: "/api/getTarotCardsList", Handler: controllers.GetAllTarotCardsHandler(db), Method: "GET"},
		{Path: "/api/getPlaceholderDeck", Handler: controllers.GetPlaceholderDeckHandler(db), Method: "GET"},

		// Three Card Reading Routes
		{Path: "/api/three-card-reading", Handler: controllers.GenerateThreeCardReading(db), Method: "POST"},
		{Path: "/api/three-card-readings", Handler: controllers.GetAllReadings(db), Method: "GET"},
		{Path: "/api/three-card-readings/{id}", Handler: controllers.GetReadingByID(db), Method: "GET"},

		// Love Tarot Reading Routes
		{Path: "/api/love-tarot-reading", Handler: controllers.GenerateLoveTarotReading(db), Method: "POST"},
		{Path: "/api/love-tarot-readings", Handler: controllers.GetAllLoveReadings(db), Method: "GET"},
		{Path: "/api/love-tarot-readings/{id}", Handler: controllers.GetLoveReadingByID(db), Method: "GET"},

		// Success Reading Routes
		{Path: "/api/success-reading", Handler: controllers.GenerateSuccessReading(db), Method: "POST"},
		{Path: "/api/success-readings", Handler: controllers.GetAllSuccessReadings(db), Method: "GET"},
		{Path: "/api/success-readings/{id}", Handler: controllers.GetSuccessReadingByID(db), Method: "GET"},

		// Spiritual Guidance Reading Routes
		{Path: "/api/spiritual-guidance-reading", Handler: controllers.GenerateSpiritualGuidanceReading(db), Method: "POST"},
		{Path: "/api/spiritual-guidance-readings", Handler: controllers.GetAllSpiritualGuidanceReadings(db), Method: "GET"},
		{Path: "/api/spiritual-guidance-readings/{id}", Handler: controllers.GetSpiritualGuidanceReadingByID(db), Method: "GET"},
	}
}

// LoadRoutes sets up the routes in the router with the provided database connection
func LoadRoutes(router *mux.Router, db *gorm.DB) {
	routes := GetRoutes(db)

	for _, route := range routes {
		router.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	}
}
