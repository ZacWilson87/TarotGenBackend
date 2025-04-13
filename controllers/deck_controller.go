package controllers

import (
	"backend_tarot/models"
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

func GetAllDecks(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var decks []models.Deck
		db.Find(&decks) //todo: add filter for userID when auth is added
		json.NewEncoder(w).Encode(decks)
	}
}

func GetDeckByID(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		deckID := r.URL.Query().Get("id")
		var deck models.Deck
		db.First(&deck, deckID)
		json.NewEncoder(w).Encode(deck)
	}
}
