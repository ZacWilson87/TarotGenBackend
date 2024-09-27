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

// GenerateSuccessReading handles the POST request to generate and save a new success reading
func GenerateSuccessReading(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Generate a new success reading
		reading, err := handlers.GenerateAndSaveSuccessReading(db)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to generate success reading: %v", err), http.StatusInternalServerError)
			return
		}

		// Return the reading as JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(reading)
	}
}

// GetAllSuccessReadings handles the GET request to retrieve all success readings
func GetAllSuccessReadings(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		readings, err := models.GetAllSuccessReadings(db)
		if err != nil {
			http.Error(w, "Failed to retrieve success readings", http.StatusInternalServerError)
			return
		}

		// Set the IsReversed flags on the associated cards
		for i := range readings {
			readings[i].FirstCard.IsReversed = readings[i].FirstCardIsReversed
			readings[i].SecondCard.IsReversed = readings[i].SecondCardIsReversed
			readings[i].ThirdCard.IsReversed = readings[i].ThirdCardIsReversed
			readings[i].FourthCard.IsReversed = readings[i].FourthCardIsReversed
			readings[i].FifthCard.IsReversed = readings[i].FifthCardIsReversed
		}

		// Return the readings as JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(readings)
	}
}

// GetSuccessReadingByID handles the GET request to retrieve a success reading by its ID
func GetSuccessReadingByID(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idParam := vars["id"]
		id, err := strconv.ParseUint(idParam, 10, 32)
		if err != nil {
			http.Error(w, "Invalid reading ID", http.StatusBadRequest)
			return
		}

		reading, err := models.GetSuccessReadingByID(db, uint(id))
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				http.Error(w, "Success reading not found", http.StatusNotFound)
			} else {
				http.Error(w, "Failed to retrieve success reading", http.StatusInternalServerError)
			}
			return
		}

		// Set the IsReversed flags on the associated cards
		reading.FirstCard.IsReversed = reading.FirstCardIsReversed
		reading.SecondCard.IsReversed = reading.SecondCardIsReversed
		reading.ThirdCard.IsReversed = reading.ThirdCardIsReversed
		reading.FourthCard.IsReversed = reading.FourthCardIsReversed
		reading.FifthCard.IsReversed = reading.FifthCardIsReversed

		// Return the reading as JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(reading)
	}
}
