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

// GenerateLoveTarotReading handles the POST request to generate and save a new love reading
func GenerateLoveTarotReading(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Generate a new love reading
		reading, err := handlers.GenerateAndSaveLoveTarotReading(db)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to generate love reading: %v", err), http.StatusInternalServerError)
			return
		}

		// Return the reading as JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(reading)
	}
}

// GetAllLoveReadings handles the GET request to retrieve all love readings
func GetAllLoveReadings(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		readings, err := models.GetAllLoveTarotReadings(db)
		if err != nil {
			http.Error(w, "Failed to retrieve love readings", http.StatusInternalServerError)
			return
		}

		// Set the IsReversed flags on the associated cards
		for i := range readings {
			readings[i].FirstCard.IsReversed = readings[i].FirstCardIsReversed
			readings[i].SecondCard.IsReversed = readings[i].SecondCardIsReversed
			readings[i].ThirdCard.IsReversed = readings[i].ThirdCardIsReversed
			readings[i].FourthCard.IsReversed = readings[i].FourthCardIsReversed
			readings[i].FifthCard.IsReversed = readings[i].FifthCardIsReversed
			readings[i].SixthCard.IsReversed = readings[i].SixthCardIsReversed
		}

		// Return the readings as JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(readings)
	}
}

// GetLoveReadingByID handles the GET request to retrieve a love reading by its ID
func GetLoveReadingByID(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idParam := vars["id"]
		id, err := strconv.ParseUint(idParam, 10, 32)
		if err != nil {
			http.Error(w, "Invalid reading ID", http.StatusBadRequest)
			return
		}

		reading, err := models.GetLoveTarotReadingByID(db, uint(id))
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				http.Error(w, "Love reading not found", http.StatusNotFound)
			} else {
				http.Error(w, "Failed to retrieve love reading", http.StatusInternalServerError)
			}
			return
		}

		// Set the IsReversed flags on the associated cards
		reading.FirstCard.IsReversed = reading.FirstCardIsReversed
		reading.SecondCard.IsReversed = reading.SecondCardIsReversed
		reading.ThirdCard.IsReversed = reading.ThirdCardIsReversed
		reading.FourthCard.IsReversed = reading.FourthCardIsReversed
		reading.FifthCard.IsReversed = reading.FifthCardIsReversed
		reading.SixthCard.IsReversed = reading.SixthCardIsReversed

		// Return the reading as JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(reading)
	}
}
