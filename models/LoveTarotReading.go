package models

import (
	"time"

	"gorm.io/gorm"
)

type LoveTarotReading struct {
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

	SixthCardID         uint      `json:"sixth_card_id"`
	SixthCard           TarotCard `gorm:"foreignKey:SixthCardID" json:"sixth_card"`
	SixthCardIsReversed bool      `json:"sixth_card_is_reversed"`
}

// Create saves the LoveTarotReading to the database
func (reading *LoveTarotReading) Create(db *gorm.DB) error {
	return db.Create(reading).Error
}

// GetLoveTarotReadingByID retrieves a reading by its ID
func GetLoveTarotReadingByID(db *gorm.DB, id uint) (LoveTarotReading, error) {
	var reading LoveTarotReading
	err := db.Preload("FirstCard").
		Preload("SecondCard").
		Preload("ThirdCard").
		Preload("FourthCard").
		Preload("FifthCard").
		Preload("SixthCard").
		First(&reading, id).Error
	return reading, err
}

// GetAllLoveTarotReadings retrieves all love readings
func GetAllLoveTarotReadings(db *gorm.DB) ([]LoveTarotReading, error) {
	var readings []LoveTarotReading
	err := db.Preload("FirstCard").
		Preload("SecondCard").
		Preload("ThirdCard").
		Preload("FourthCard").
		Preload("FifthCard").
		Preload("SixthCard").
		Find(&readings).Error
	return readings, err
}
