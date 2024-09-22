package seed

import (
	"backend_tarot/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"gorm.io/gorm"
)

// TarotCardSeedRoot is used to map the root JSON structure containing cards
type TarotCardSeedRoot struct {
	Cards []models.TarotCard `json:"cards"`
}

// SeedDatabase seeds the Tarot cards from seed.json into the database
func SeedDatabase(db *gorm.DB) {
	// Check if any Tarot cards already exist in the database
	var count int64
	db.Model(&models.TarotCard{}).Count(&count)

	// If records already exist, skip the seeding
	if count > 0 {
		log.Println("Database already seeded, skipping seeding process.")
		return
	}

	// Read the seed.json file
	file, err := os.Open("seed.json")
	if err != nil {
		log.Fatalf("Error opening seed.json file: %v", err)
	}
	defer file.Close()

	// Read the file contents
	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading seed.json file: %v", err)
	}

	// Parse the JSON into the TarotCardSeedRoot struct
	var seedData TarotCardSeedRoot
	err = json.Unmarshal(bytes, &seedData)
	if err != nil {
		log.Fatalf("Error unmarshalling seed.json: %v", err)
	}

	// Create a new deck
	deck := models.Deck{
		Name: "Default Tarot Deck", // You can give it a custom name if you want
	}
	if err := db.Create(&deck).Error; err != nil {
		log.Fatalf("Error creating deck: %v", err)
	}

	// Seed Tarot cards into the database
	for _, seedCard := range seedData.Cards {
		// Set the DeckID for each card to the newly created deck's ID
		seedCard.DeckID = deck.ID

		// Insert Tarot card into the database
		if err := db.Create(&seedCard).Error; err != nil {
			log.Printf("Error seeding Tarot card %s: %v", seedCard.Name, err)
		}
	}

	fmt.Println("Database seeded successfully!")
}
