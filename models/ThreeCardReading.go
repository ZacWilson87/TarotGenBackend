package models

import (
	"time"

	"golang.org/x/exp/rand"
	"gorm.io/gorm"
)

type ThreeCardReading struct {
	PastCard    string `json:"past_card"`
	PresentCard string `json:"present_card"`
	FutureCard  string `json:"future_card"`
	Reading     string `json:"explanation"`
	Date        string `json:"date"`
}

func shuffleDeck(db *gorm.DB) []TarotCard {
	// get placeholder deck
	deck, err := GetPlaceholderDeck(db)
	if err != nil {
		panic(err)
	}
	// shuffle deck
	rand.Seed(uint64(time.Now().UnixNano()))
	rand.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })
	deck = reverseRandomCards(deck)
	return deck
}

func getIsReversed() bool {
	// 15% chance of returning true
	return rand.Intn(100) < 15
}

func reverseRandomCards(deck []TarotCard) []TarotCard {
	// 15% chance of reversing cards
	if getIsReversed() {
		rand.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })
	}
	return deck
}
