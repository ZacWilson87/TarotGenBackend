package models

import (
	"time"

	"gorm.io/gorm"
)

type SpiritualGuidanceReading struct {
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

	SeventhCardID         uint      `json:"seventh_card_id"`
	SeventhCard           TarotCard `gorm:"foreignKey:SeventhCardID" json:"seventh_card"`
	SeventhCardIsReversed bool      `json:"seventh_card_is_reversed"`

	EighthCardID         uint      `json:"eighth_card_id"`
	EighthCard           TarotCard `gorm:"foreignKey:EighthCardID" json:"eighth_card"`
	EighthCardIsReversed bool      `json:"eighth_card_is_reversed"`
}

// Create saves the SpiritualGuidanceReading to the database
func (reading *SpiritualGuidanceReading) Create(db *gorm.DB) error {
	return db.Create(reading).Error
}

// GetSpiritualGuidanceReadingByID retrieves a reading by its ID
func GetSpiritualGuidanceReadingByID(db *gorm.DB, id uint) (SpiritualGuidanceReading, error) {
	var reading SpiritualGuidanceReading
	err := db.Preload("FirstCard").
		Preload("SecondCard").
		Preload("ThirdCard").
		Preload("FourthCard").
		Preload("FifthCard").
		Preload("SixthCard").
		Preload("SeventhCard").
		Preload("EighthCard").
		First(&reading, id).Error
	return reading, err
}

// GetAllSpiritualGuidanceReadings retrieves all spiritual guidance readings
func GetAllSpiritualGuidanceReadings(db *gorm.DB) ([]SpiritualGuidanceReading, error) {
	var readings []SpiritualGuidanceReading
	err := db.Preload("FirstCard").
		Preload("SecondCard").
		Preload("ThirdCard").
		Preload("FourthCard").
		Preload("FifthCard").
		Preload("SixthCard").
		Preload("SeventhCard").
		Preload("EighthCard").
		Find(&readings).Error
	return readings, err
}

// DeleteSpiritualGuidanceReadingByID deletes a reading by its ID
func DeleteSpiritualGuidanceReadingByID(db *gorm.DB, id uint) error {
	return db.Delete(&SpiritualGuidanceReading{}, id).Error
}
