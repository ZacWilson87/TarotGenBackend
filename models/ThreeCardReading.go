// models/ThreeCardReading.go

package models

import (
	"time"

	"gorm.io/gorm"
)

type ThreeCardReading struct {
	ID                    uint      `gorm:"primaryKey" json:"id"`
	PastCardID            uint      `json:"past_card_id"`
	PastCard              TarotCard `gorm:"foreignKey:PastCardID" json:"past_card"`
	PastCardIsReversed    bool      `gorm:"not null" json:"past_card_is_reversed"`
	PresentCardID         uint      `json:"present_card_id"`
	PresentCard           TarotCard `gorm:"foreignKey:PresentCardID" json:"present_card"`
	PresentCardIsReversed bool      `gorm:"not null" json:"present_card_is_reversed"`
	FutureCardID          uint      `json:"future_card_id"`
	FutureCard            TarotCard `gorm:"foreignKey:FutureCardID" json:"future_card"`
	FutureCardIsReversed  bool      `gorm:"not null" json:"future_card_is_reversed"`
	Reading               string    `gorm:"type:text" json:"reading"`
	Date                  time.Time `json:"date"`
}

// Create saves the ThreeCardReading to the database
func (reading *ThreeCardReading) Create(db *gorm.DB) error {
	return db.Create(reading).Error
}

// GetThreeCardReadingByID retrieves a reading by its ID, including associated cards
func GetThreeCardReadingByID(db *gorm.DB, id uint) (ThreeCardReading, error) {
	var reading ThreeCardReading
	err := db.Preload("PastCard").
		Preload("PresentCard").
		Preload("FutureCard").
		First(&reading, id).Error
	return reading, err
}

// GetAllThreeCardReadings retrieves all readings from the database
func GetAllThreeCardReadings(db *gorm.DB) ([]ThreeCardReading, error) {
	var readings []ThreeCardReading
	err := db.Preload("PastCard").
		Preload("PresentCard").
		Preload("FutureCard").
		Find(&readings).Error
	return readings, err
}
