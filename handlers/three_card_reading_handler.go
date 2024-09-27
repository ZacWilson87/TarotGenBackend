// handlers/three_card_reading_handler.go

package handlers

import (
	"backend_tarot/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"gorm.io/gorm"
)

// GenerateThreeCardReadingHandler handles generating and storing a three-card reading
func GenerateThreeCardReadingHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Shuffle the deck
		deck, err := models.ShuffleDeck(db)
		if err != nil {
			http.Error(w, "Failed to shuffle deck", http.StatusInternalServerError)
			return
		}

		if len(deck) < 3 {
			http.Error(w, "Not enough cards in the deck", http.StatusInternalServerError)
			return
		}

		// Select the top three cards
		pastCard := deck[0]
		presentCard := deck[1]
		futureCard := deck[2]

		// Generate reading explanation using OpenAI
		explanation, err := GenerateReadingExplanation(pastCard, presentCard, futureCard)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error generating reading: %v", err), http.StatusInternalServerError)
			return
		}

		// Create a ThreeCardReading instance
		reading := models.ThreeCardReading{
			PastCardID:            pastCard.ID,
			PastCardIsReversed:    pastCard.IsReversed,
			PresentCardID:         presentCard.ID,
			PresentCardIsReversed: presentCard.IsReversed,
			FutureCardID:          futureCard.ID,
			FutureCardIsReversed:  futureCard.IsReversed,
			Reading:               explanation,
			Date:                  time.Now(),
		}

		// Save the reading to the database
		err = reading.Create(db)
		if err != nil {
			http.Error(w, "Failed to save reading", http.StatusInternalServerError)
			return
		}

		// Retrieve the reading with associated cards
		readingWithCards, err := models.GetThreeCardReadingByID(db, reading.ID)
		if err != nil {
			http.Error(w, "Failed to retrieve reading", http.StatusInternalServerError)
			return
		}

		// Set the IsReversed flags on the associated cards
		readingWithCards.PastCard.IsReversed = readingWithCards.PastCardIsReversed
		readingWithCards.PresentCard.IsReversed = readingWithCards.PresentCardIsReversed
		readingWithCards.FutureCard.IsReversed = readingWithCards.FutureCardIsReversed

		// Return the reading as JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(readingWithCards)
	}
}

func GenerateAndSaveThreeCardReading(db *gorm.DB) (models.ThreeCardReading, error) {
	// Shuffle the deck
	deck, err := models.ShuffleDeck(db)
	if err != nil {
		return models.ThreeCardReading{}, fmt.Errorf("failed to shuffle deck: %v", err)
	}

	if len(deck) < 3 {
		return models.ThreeCardReading{}, fmt.Errorf("not enough cards in the deck")
	}

	// Select the top three cards
	pastCard := deck[0]
	presentCard := deck[1]
	futureCard := deck[2]

	// Generate reading explanation using OpenAI
	explanation, err := GenerateReadingExplanation(pastCard, presentCard, futureCard)
	if err != nil {
		return models.ThreeCardReading{}, fmt.Errorf("error generating reading: %v", err)
	}

	// Create a ThreeCardReading instance
	reading := models.ThreeCardReading{
		PastCardID:            pastCard.ID,
		PastCardIsReversed:    pastCard.IsReversed,
		PresentCardID:         presentCard.ID,
		PresentCardIsReversed: presentCard.IsReversed,
		FutureCardID:          futureCard.ID,
		FutureCardIsReversed:  futureCard.IsReversed,
		Reading:               explanation,
		Date:                  time.Now(),
	}

	// Save the reading to the database
	err = reading.Create(db)
	if err != nil {
		return models.ThreeCardReading{}, fmt.Errorf("failed to save reading: %v", err)
	}

	// Retrieve the reading with associated cards
	readingWithCards, err := models.GetThreeCardReadingByID(db, reading.ID)
	if err != nil {
		return models.ThreeCardReading{}, fmt.Errorf("failed to retrieve reading: %v", err)
	}

	// Set the IsReversed flags on the associated cards
	readingWithCards.PastCard.IsReversed = readingWithCards.PastCardIsReversed
	readingWithCards.PresentCard.IsReversed = readingWithCards.PresentCardIsReversed
	readingWithCards.FutureCard.IsReversed = readingWithCards.FutureCardIsReversed

	return readingWithCards, nil
}

func GenerateReadingExplanation(past models.TarotCard, present models.TarotCard, future models.TarotCard) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	// Construct the prompt for OpenAI
	prompt := fmt.Sprintf(
		"Provide a detailed three-card tarot reading interpretation.\n\n"+
			"Past Card:\nName: %s\nMeaning: %s\nIs Reversed: %t\n\n"+
			"Present Card:\nName: %s\nMeaning: %s\nIs Reversed: %t\n\n"+
			"Future Card:\nName: %s\nMeaning: %s\nIs Reversed: %t\n\n"+
			"Interpretation:",
		past.Name, getMeaning(past), past.IsReversed,
		present.Name, getMeaning(present), present.IsReversed,
		future.Name, getMeaning(future), future.IsReversed,
	)

	// Prepare the request payload
	requestBody, err := json.Marshal(map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []map[string]string{
			{"role": "system", "content": "You are a helpful tarot card interpreter."},
			{"role": "user", "content": prompt},
		},
		"max_tokens":  500,
		"n":           1,
		"temperature": 0.7,
	})
	if err != nil {
		return "", err
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read and parse the response
	var openAIResponse models.OpenAIResponse
	err = json.NewDecoder(resp.Body).Decode(&openAIResponse)
	if err != nil {
		return "", err
	}

	if len(openAIResponse.Choices) == 0 {
		return "", fmt.Errorf("no choices returned from OpenAI")
	}

	// Return the generated explanation
	return openAIResponse.Choices[0].Message.Content, nil
}

func getMeaning(card models.TarotCard) string {
	if card.IsReversed {
		return card.ReversedMeaning
	}
	return card.Meaning
}
