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

func GenerateAndSaveSpiritualGuidanceReading(db *gorm.DB) (models.SpiritualGuidanceReading, error) {
	// Shuffle the deck
	deck, err := models.ShuffleDeck(db)
	if err != nil {
		return models.SpiritualGuidanceReading{}, fmt.Errorf("failed to shuffle deck: %v", err)
	}

	if len(deck) < 8 {
		return models.SpiritualGuidanceReading{}, fmt.Errorf("not enough cards in the deck")
	}

	// Select the top eight cards
	firstCard := deck[0]
	secondCard := deck[1]
	thirdCard := deck[2]
	fourthCard := deck[3]
	fifthCard := deck[4]
	sixthCard := deck[5]
	seventhCard := deck[6]
	eighthCard := deck[7]

	// Generate reading explanation using OpenAI
	explanation, err := GenerateSpiritualGuidanceReadingExplanation(
		firstCard, secondCard, thirdCard, fourthCard, fifthCard, sixthCard, seventhCard, eighthCard)
	if err != nil {
		return models.SpiritualGuidanceReading{}, fmt.Errorf("error generating reading: %v", err)
	}

	// Create a SpiritualGuidanceReading instance
	reading := models.SpiritualGuidanceReading{
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

		SeventhCardID:         seventhCard.ID,
		SeventhCardIsReversed: seventhCard.IsReversed,

		EighthCardID:         eighthCard.ID,
		EighthCardIsReversed: eighthCard.IsReversed,
	}

	// Save the reading to the database
	err = reading.Create(db)
	if err != nil {
		return models.SpiritualGuidanceReading{}, fmt.Errorf("failed to save reading: %v", err)
	}

	// Retrieve the reading with associated cards
	readingWithCards, err := models.GetSpiritualGuidanceReadingByID(db, reading.ID)
	if err != nil {
		return models.SpiritualGuidanceReading{}, fmt.Errorf("failed to retrieve reading: %v", err)
	}

	// Set the IsReversed flags on the associated cards
	readingWithCards.FirstCard.IsReversed = readingWithCards.FirstCardIsReversed
	readingWithCards.SecondCard.IsReversed = readingWithCards.SecondCardIsReversed
	readingWithCards.ThirdCard.IsReversed = readingWithCards.ThirdCardIsReversed
	readingWithCards.FourthCard.IsReversed = readingWithCards.FourthCardIsReversed
	readingWithCards.FifthCard.IsReversed = readingWithCards.FifthCardIsReversed
	readingWithCards.SixthCard.IsReversed = readingWithCards.SixthCardIsReversed
	readingWithCards.SeventhCard.IsReversed = readingWithCards.SeventhCardIsReversed
	readingWithCards.EighthCard.IsReversed = readingWithCards.EighthCardIsReversed

	return readingWithCards, nil
}

func GenerateSpiritualGuidanceReadingExplanation(
	firstCard, secondCard, thirdCard, fourthCard, fifthCard, sixthCard, seventhCard, eighthCard models.TarotCard,
) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	// Construct the prompt for OpenAI
	prompt := fmt.Sprintf(
		"Provide a detailed eight-card spiritual guidance tarot reading interpretation.\n\n"+
			"#1 – Your main concerns, going in-depth into the problem:\n"+
			"Card Name: %s\nMeaning: %s\nIs Reversed: %t\n\n"+
			"#2 – Your motivation for seeking guidance:\n"+
			"Card Name: %s\nMeaning: %s\nIs Reversed: %t\n\n"+
			"#3 – Things in your life you're insecure or worried about:\n"+
			"Card Name: %s\nMeaning: %s\nIs Reversed: %t\n\n"+
			"#4 – Parts of your life that you are not aware of:\n"+
			"Card Name: %s\nMeaning: %s\nIs Reversed: %t\n\n"+
			"#5 – Advice to guide you to face your fears, tying in with previous cards:\n"+
			"Card Name: %s\nMeaning: %s\nIs Reversed: %t\n\n"+
			"#6 – Guidance to a life with no worries to move forward on your spiritual journey:\n"+
			"Card Name: %s\nMeaning: %s\nIs Reversed: %t\n\n"+
			"#7 – Teaches you to deal with the situation with resources at hand:\n"+
			"Card Name: %s\nMeaning: %s\nIs Reversed: %t\n\n"+
			"#8 – Concludes by telling that the result depends on your reaction, focusing on positive or negative:\n"+
			"Card Name: %s\nMeaning: %s\nIs Reversed: %t\n\n"+
			"Interpretation:",
		firstCard.Name, getMeaning(firstCard), firstCard.IsReversed,
		secondCard.Name, getMeaning(secondCard), secondCard.IsReversed,
		thirdCard.Name, getMeaning(thirdCard), thirdCard.IsReversed,
		fourthCard.Name, getMeaning(fourthCard), fourthCard.IsReversed,
		fifthCard.Name, getMeaning(fifthCard), fifthCard.IsReversed,
		sixthCard.Name, getMeaning(sixthCard), sixthCard.IsReversed,
		seventhCard.Name, getMeaning(seventhCard), seventhCard.IsReversed,
		eighthCard.Name, getMeaning(eighthCard), eighthCard.IsReversed,
	)

	// Prepare the request payload
	requestBody, err := json.Marshal(map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []map[string]string{
			{"role": "system", "content": "You are a helpful tarot card interpreter specializing in spiritual guidance readings."},
			{"role": "user", "content": prompt},
		},
		"max_tokens":  1500,
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
