package controllers

import (
	"backend_tarot/handlers"
	"backend_tarot/models"
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

// TarotCardRequest is the expected structure of the incoming request for generating a tarot card
type TarotCardRequest struct {
	Card   string `json:"tarotCard"`
	Theme  string `json:"theme"`
	Color1 string `json:"color1"`
}

func GenerateTarotCardHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode the incoming JSON request
		var req TarotCardRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// Validate the Tarot card
		if !models.IsValidTarotCard(db, req.Card) {
			http.Error(w, "Invalid tarot card", http.StatusBadRequest)
			return
		}

		// find card via name and return card description
		var tarotCard models.TarotCard
		err = db.Where("name = ?", req.Card).First(&tarotCard).Error
		if err != nil {
			http.Error(w, "Tarot card not found", http.StatusNotFound)
			return
		}

		// Generate Tarot card design using OpenAI API
		cardDesignObj, err := handlers.GenerateTarotCardDesign(tarotCard.Description, req.Theme, req.Color1)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error generating tarot card: %v", err), http.StatusInternalServerError)
			return
		}

		// Download and save the image locally
		imagePath, err := handlers.DownloadImage(cardDesignObj.Url, req.Card)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error downloading tarot card image: %v", err), http.StatusInternalServerError)
			return
		}

		// Retrieve the deck
		var deck models.Deck
		err = db.First(&deck).Error // Replace with appropriate deck fetching logic
		if err != nil {
			http.Error(w, "Deck not found", http.StatusInternalServerError)
			return
		}

		// Create a new TarotCard object
		newCard := models.TarotCard{
			Name:            req.Card,
			ArcanaType:      tarotCard.ArcanaType, // Logic to determine Major/Minor Arcana
			FilePath:        imagePath,
			Meaning:         tarotCard.Meaning,         // Use the correct logic to set the meaning
			ReversedMeaning: tarotCard.ReversedMeaning, // Use the correct logic to set the reversed meaning
			Description:     cardDesignObj.Description,
			Suit:            tarotCard.Suit,
		}

		// Add the card to the deck (or create a new deck if it already exists)
		err = models.AddTarotCard(db, &deck, newCard)
		if err != nil {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}

		// Return success response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Tarot card added to deck",
			"design":  newCard.FilePath,
		})
	}
}

// GetAllTarotCardsHandler handles the /tarotCards GET request
// It retrieves all tarot cards from the database
func GetAllTarotCardsHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Retrieve all tarot cards from the database
		cards, err := models.GetAllTarotCards(db)
		if err != nil {
			http.Error(w, "Failed to retrieve tarot cards", http.StatusInternalServerError)
			return
		}

		// Return the tarot cards as JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cards)
	}
}

// GetPlaceholderDeckHandler handles the /placeholderDeck GET request
// It retrieves the placeholder deck from the database
func GetPlaceholderDeckHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Retrieve the placeholder deck from the database
		deck, err := models.GetPlaceholderDeck(db)
		if err != nil {
			http.Error(w, "Failed to retrieve placeholder deck", http.StatusInternalServerError)
			return
		}

		// Return the placeholder deck as JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(deck)
	}
}
