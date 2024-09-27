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

// GenerateSpiritualGuidanceReading handles the POST request to generate and save a new spiritual guidance reading
func GenerateSpiritualGuidanceReading(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Generate a new spiritual guidance reading
		reading, err := handlers.GenerateAndSaveSpiritualGuidanceReading(db)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to generate spiritual guidance reading: %v", err), http.StatusInternalServerError)
			return
		}

		// Return the reading as JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(reading)
	}
}

// GetAllSpiritualGuidanceReadings handles the GET request to retrieve all spiritual guidance readings
func GetAllSpiritualGuidanceReadings(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		readings, err := models.GetAllSpiritualGuidanceReadings(db)
		if err != nil {
			http.Error(w, "Failed to retrieve spiritual guidance readings", http.StatusInternalServerError)
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
			readings[i].SeventhCard.IsReversed = readings[i].SeventhCardIsReversed
			readings[i].EighthCard.IsReversed = readings[i].EighthCardIsReversed
		}

		// Return the readings as JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(readings)
	}
}

// GetSpiritualGuidanceReadingByID handles the GET request to retrieve a spiritual guidance reading by its ID
func GetSpiritualGuidanceReadingByID(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idParam := vars["id"]
		id, err := strconv.ParseUint(idParam, 10, 32)
		if err != nil {
			http.Error(w, "Invalid reading ID", http.StatusBadRequest)
			return
		}

		reading, err := models.GetSpiritualGuidanceReadingByID(db, uint(id))
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				http.Error(w, "Spiritual guidance reading not found", http.StatusNotFound)
			} else {
				http.Error(w, "Failed to retrieve spiritual guidance reading", http.StatusInternalServerError)
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
		reading.SeventhCard.IsReversed = reading.SeventhCardIsReversed
		reading.EighthCard.IsReversed = reading.EighthCardIsReversed

		// Return the reading as JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(reading)
	}
}
