// controllers/three_card_reading_controller.go

package controllers

import (
	"backend_tarot/handlers"
	"backend_tarot/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// GenerateThreeCardReading handles the POST request to generate and save a new reading
func GenerateThreeCardReading(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Generate a new reading using the handler
		reading, err := handlers.GenerateAndSaveThreeCardReading(db)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to generate reading: %v", err), http.StatusInternalServerError)
			return
		}

		// Return the reading as JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(reading)
	}
}

// GetAllReadings handles the GET request to retrieve all readings
func GetAllReadings(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		readings, err := models.GetAllThreeCardReadings(db)
		if err != nil {
			http.Error(w, "Failed to retrieve readings", http.StatusInternalServerError)
			return
		}

		// Return the readings as JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(readings)
	}
}

// GetReadingByID handles the GET request to retrieve a reading by its ID
func GetReadingByID(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idParam := vars["id"]
		id, err := strconv.ParseUint(idParam, 10, 32)
		if err != nil {
			http.Error(w, "Invalid reading ID", http.StatusBadRequest)
			return
		}

		reading, err := models.GetThreeCardReadingByID(db, uint(id))
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				http.Error(w, "Reading not found", http.StatusNotFound)
			} else {
				http.Error(w, "Failed to retrieve reading", http.StatusInternalServerError)
			}
			return
		}

		// Set the IsReversed flags on the associated cards
		reading.PastCard.IsReversed = reading.PastCardIsReversed
		reading.PresentCard.IsReversed = reading.PresentCardIsReversed
		reading.FutureCard.IsReversed = reading.FutureCardIsReversed

		// Return the reading as JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(reading)
	}
}
