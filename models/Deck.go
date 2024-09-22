package models

import (
	"fmt"

	"gorm.io/gorm"
)

// Deck defines the structure for a Tarot deck
type Deck struct {
	ID                uint        `gorm:"primaryKey"`        // Auto-incrementing ID for GORM
	UserID            uint        `gorm:"not null"`          // Foreign key to the User (if applicable)
	Name              string      `gorm:"size:100;not null"` // Deck name (e.g., "My Custom Deck")
	Cards             []TarotCard `gorm:"foreignKey:DeckID"` // One-to-many relationship: a deck has many Tarot cards
	IsPlaceholderDeck bool        `gorm:"not null"`          // Whether the deck is a placeholder deck
}

// AddTarotCard adds a tarot card to the deck if it doesn't already exist.
// If the card already exists in the deck, it creates a new deck and adds the card there.
func AddTarotCard(db *gorm.DB, deck *Deck, card TarotCard) error {
	// Check if the card already exists in the deck
	var existingCard TarotCard
	if db.Where("name = ? AND deck_id = ?", card.Name, deck.ID).First(&existingCard).RowsAffected > 0 {
		// Tarot card already exists in the deck, so create a new deck
		newDeck := Deck{
			UserID:            deck.UserID,
			Name:              fmt.Sprintf("%s's New Deck", deck.Name), // Generate a name for the new deck
			IsPlaceholderDeck: false,                                   // You can set this based on your logic
		}

		// Save the new deck
		if err := db.Create(&newDeck).Error; err != nil {
			return fmt.Errorf("failed to create a new deck: %w", err)
		}

		// Add the card to the newly created deck
		card.DeckID = newDeck.ID
		if err := db.Create(&card).Error; err != nil {
			return fmt.Errorf("failed to add tarot card to the new deck: %w", err)
		}

		return nil
	}

	// Add the card to the existing deck
	card.DeckID = deck.ID
	if err := db.Create(&card).Error; err != nil {
		return fmt.Errorf("failed to add tarot card to deck: %w", err)
	}

	return nil
}

// GetDeck retrieves a deck by ID, including its associated Tarot cards
func GetDeck(db *gorm.DB, deckID uint) (*Deck, error) {
	var deck Deck
	err := db.Preload("Cards").First(&deck, deckID).Error
	if err != nil {
		return nil, err
	}
	return &deck, nil
}
