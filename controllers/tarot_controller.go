// controllers/tarot_controller.go
package controllers

import (
	"backend_tarot/handlers"
	"backend_tarot/models"
	"encoding/json"
	"fmt"
	"net/http"
)

// TarotCardRequest is the expected structure of the incoming request
type TarotCardRequest struct {
	Card   string `json:"tarotCard"`
	Theme  string `json:"theme"`
	Color1 string `json:"color1"`
}

// GenerateTarotCardHandler handles the /generateTarotCard POST request
func GenerateTarotCardHandler(w http.ResponseWriter, r *http.Request) {
	// Decode the incoming JSON request
	var req TarotCardRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	fmt.Println(req)
	// Validate the input: Check if card is valid
	if !models.IsValidTarotCard(models.TarotCard(req.Card)) {
		http.Error(w, "Invalid tarot card", http.StatusBadRequest)
		return
	}

	// Call the handler function to generate the tarot card design
	cardDesign, err := handlers.GenerateTarotCardDesign(req.Card, req.Theme, req.Color1)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error generating tarot card: %v", err), http.StatusInternalServerError)
		return
	}

	// Return the generated card design as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"design": cardDesign,
	})
}

// GetAllTarotCardsHandler handles the /tarotCards GET request
func GetAllTarotCardsHandler(w http.ResponseWriter, r *http.Request) {
	cards := models.GetAllTarotCards()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cards)
}
