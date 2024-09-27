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

func GenerateAndSaveLoveTarotReading(db *gorm.DB) (models.LoveTarotReading, error) {
	// Shuffle the deck
	deck, err := models.ShuffleDeck(db)
	if err != nil {
		return models.LoveTarotReading{}, fmt.Errorf("failed to shuffle deck: %v", err)
	}

	if len(deck) < 6 {
		return models.LoveTarotReading{}, fmt.Errorf("not enough cards in the deck")
	}

	// Select the top six cards
	firstCard := deck[0]
	secondCard := deck[1]
	thirdCard := deck[2]
	fourthCard := deck[3]
	fifthCard := deck[4]
	sixthCard := deck[5]

	// Generate reading explanation using OpenAI
	explanation, err := GenerateLoveReadingExplanation(firstCard, secondCard, thirdCard, fourthCard, fifthCard, sixthCard)
	if err != nil {
		return models.LoveTarotReading{}, fmt.Errorf("error generating reading: %v", err)
	}

	// Create a LoveTarotReading instance
	reading := models.LoveTarotReading{
		Date:    time.Now(),
		Reading: explanation,

		FirstCardID:         firstCard.ID,
		FirstCardIsReversed: firstCard.IsReversed,

		SecondCardID:         secondCard.ID,
		SecondCardIsReversed: secondCard.IsReversed,

		ThirdCardID:         thirdCard.ID,
		ThirdCardIsReversed: thirdCard.IsReversed,

		FourthCardID:         fourthCard.ID,
		FourthCardIsReversed: fourthCard.IsReversed,

		FifthCardID:         fifthCard.ID,
		FifthCardIsReversed: fifthCard.IsReversed,

		SixthCardID:         sixthCard.ID,
		SixthCardIsReversed: sixthCard.IsReversed,
	}

	// Save the reading to the database
	err = reading.Create(db)
	if err != nil {
		return models.LoveTarotReading{}, fmt.Errorf("failed to save reading: %v", err)
	}

	// Retrieve the reading with associated cards
	readingWithCards, err := models.GetLoveTarotReadingByID(db, reading.ID)
	if err != nil {
		return models.LoveTarotReading{}, fmt.Errorf("failed to retrieve reading: %v", err)
	}

	// Set the IsReversed flags on the associated cards
	readingWithCards.FirstCard.IsReversed = readingWithCards.FirstCardIsReversed
	readingWithCards.SecondCard.IsReversed = readingWithCards.SecondCardIsReversed
	readingWithCards.ThirdCard.IsReversed = readingWithCards.ThirdCardIsReversed
	readingWithCards.FourthCard.IsReversed = readingWithCards.FourthCardIsReversed
	readingWithCards.FifthCard.IsReversed = readingWithCards.FifthCardIsReversed
	readingWithCards.SixthCard.IsReversed = readingWithCards.SixthCardIsReversed

	return readingWithCards, nil
}

func GenerateLoveReadingExplanation(
	firstCard, secondCard, thirdCard, fourthCard, fifthCard, sixthCard models.TarotCard,
) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	// Construct the prompt for OpenAI
	prompt := fmt.Sprintf(
		"Provide a detailed six-card love tarot reading interpretation.\n\n"+
			"#1 - Your current feelings about your relationship, your approach, and your outlook:\n"+
			"Card Name: %s\nMeaning: %s\nIs Reversed: %t\n\n"+
			"#2 - Your partner's current emotions towards you, their attitude, and expectations about your relationship:\n"+
			"Card Name: %s\nMeaning: %s\nIs Reversed: %t\n\n"+
			"#3 - Connection card (common characteristics of both of you):\n"+
			"Card Name: %s\nMeaning: %s\nIs Reversed: %t\n\n"+
			"#4 - The strength of your relationship:\n"+
			"Card Name: %s\nMeaning: %s\nIs Reversed: %t\n\n"+
			"#5 - The weaknesses in your relationship:\n"+
			"Card Name: %s\nMeaning: %s\nIs Reversed: %t\n\n"+
			"#6 - Your true love card (interprets if the relationship is going to be successful or not):\n"+
			"Card Name: %s\nMeaning: %s\nIs Reversed: %t\n\n"+
			"Interpretation:",
		firstCard.Name, getMeaning(firstCard), firstCard.IsReversed,
		secondCard.Name, getMeaning(secondCard), secondCard.IsReversed,
		thirdCard.Name, getMeaning(thirdCard), thirdCard.IsReversed,
		fourthCard.Name, getMeaning(fourthCard), fourthCard.IsReversed,
		fifthCard.Name, getMeaning(fifthCard), fifthCard.IsReversed,
		sixthCard.Name, getMeaning(sixthCard), sixthCard.IsReversed,
	)

	// Prepare the request payload
	requestBody, err := json.Marshal(map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []map[string]string{
			{"role": "system", "content": "You are a helpful tarot card interpreter specializing in love readings."},
			{"role": "user", "content": prompt},
		},
		"max_tokens":  1000,
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
