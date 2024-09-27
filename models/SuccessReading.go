package models

import (
	"time"

	"gorm.io/gorm"
)

type SuccessReading struct {
	ID      uint      `gorm:"primaryKey" json:"id"`
	Date    time.Time `json:"date"`
	Reading string    `gorm:"type:text" json:"reading"`

	FirstCardID         uint      `json:"first_card_id"`
	FirstCard           TarotCard `gorm:"foreignKey:FirstCardID" json:"first_card"`
	FirstCardIsReversed bool      `json:"first_card_is_reversed"`

	SecondCardID         uint      `json:"second_card_id"`
	SecondCard           TarotCard `gorm:"foreignKey:SecondCardID" json:"second_card"`
	SecondCardIsReversed bool      `json:"second_card_is_reversed"`

	ThirdCardID         uint      `json:"third_card_id"`
	ThirdCard           TarotCard `gorm:"foreignKey:ThirdCardID" json:"third_card"`
	ThirdCardIsReversed bool      `json:"third_card_is_reversed"`

	FourthCardID         uint      `json:"fourth_card_id"`
	FourthCard           TarotCard `gorm:"foreignKey:FourthCardID" json:"fourth_card"`
	FourthCardIsReversed bool      `json:"fourth_card_is_reversed"`

	FifthCardID         uint      `json:"fifth_card_id"`
	FifthCard           TarotCard `gorm:"foreignKey:FifthCardID" json:"fifth_card"`
	FifthCardIsReversed bool      `json:"fifth_card_is_reversed"`
}

// Create saves the SuccessReading to the database
func (reading *SuccessReading) Create(db *gorm.DB) error {
	return db.Create(reading).Error
}

// GetSuccessReadingByID retrieves a reading by its ID
func GetSuccessReadingByID(db *gorm.DB, id uint) (SuccessReading, error) {
	var reading SuccessReading
	err := db.Preload("FirstCard").
		Preload("SecondCard").
		Preload("ThirdCard").
		Preload("FourthCard").
		Preload("FifthCard").
		First(&reading, id).Error
	return reading, err
}

// GetAllSuccessReadings retrieves all success readings
func GetAllSuccessReadings(db *gorm.DB) ([]SuccessReading, error) {
	var readings []SuccessReading
	err := db.Preload("FirstCard").
		Preload("SecondCard").
		Preload("ThirdCard").
		Preload("FourthCard").
		Preload("FifthCard").
		Find(&readings).Error
	return readings, err
}
