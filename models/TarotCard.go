package models

import "gorm.io/gorm"

// ArcanaType defines the type of Tarot card (Major or Minor Arcana)
type ArcanaType string

const (
	MajorArcana ArcanaType = "Major"
	MinorArcana ArcanaType = "Minor"
)

// TarotCard defines the structure for the Tarot card model
type TarotCard struct {
	ID              uint       `gorm:"primaryKey" json:"id"`              // Auto-incrementing ID for GORM
	Name            string     `gorm:"size:100;not null" json:"name"`     // Tarot card name (e.g., The Fool, Ace of Wands)
	ArcanaType      ArcanaType `gorm:"size:10;not null" json:"type"`      // Major or Minor Arcana
	Suit            string     `gorm:"size:50" json:"suit,omitempty"`     // Optional field for Minor Arcana suits (e.g., Wands, Cups)
	Description     string     `gorm:"type:text" json:"description"`      // Optional field for description
	FilePath        string     `gorm:"size:255" json:"path"`              // Path to the Tarot card image
	DeckID          uint       `gorm:"not null" json:"deck_id"`           // Foreign key to the associated Deck (if applicable)
	Meaning         string     `gorm:"type:text" json:"meaning"`          // Upright Tarot card meaning
	ReversedMeaning string     `gorm:"type:text" json:"reversed_meaning"` // Reversed Tarot card meaning
	IsReversed      bool       `gorm:"not null" json:"is_reversed"`       // Whether the Tarot card is reversed
}

// IsValidTarotCard checks if a Tarot card exists in the database by name
func IsValidTarotCard(db *gorm.DB, cardName string) bool {
	var tarotCard TarotCard
	err := db.Where("name = ?", cardName).First(&tarotCard).Error
	return err == nil // If no error, the card exists
}

// GetAllTarotCards retrieves all Tarot cards from the database
func GetAllTarotCards(db *gorm.DB) ([]TarotCard, error) {
	var cards []TarotCard
	err := db.Find(&cards).Error
	if err != nil {
		return nil, err
	}
	return cards, nil
}

func GenTarotCardMetadataFromCardName(db *gorm.DB, cardName string) (TarotCard, error) {
	var tarotCard TarotCard

	err := db.Where("name = ?", cardName).First(&tarotCard).Error
	if err != nil {
		return tarotCard, err
	}
	return tarotCard, nil
}
