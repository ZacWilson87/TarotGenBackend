package main

import (
	"backend_tarot/config"
	"backend_tarot/models"
	"backend_tarot/routes"
	"backend_tarot/seed"
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Connect to the database
	config.ConnectDatabase()

	// Auto-migrate the schema for the Deck and TarotCard models
	if err := config.DB.AutoMigrate(
		&models.Deck{},
		&models.TarotCard{},
		&models.ThreeCardReading{},
		&models.LoveTarotReading{},
		&models.SuccessReading{},
		&models.SpiritualGuidanceReading{},
	); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Add a flag to trigger the database seeding process
	seedDB := flag.Bool("seed", false, "Seed the database with initial data")
	flag.Parse()

	// Check if the seed flag is set or if the database is already seeded
	if *seedDB {
		log.Println("Seeding database...")
		seed.SeedDatabase(config.DB)
	} else {
		// Seed only if the tarot_cards table is empty
		var count int64
		config.DB.Model(&models.TarotCard{}).Count(&count)
		if count == 0 {
			log.Println("Database is empty, seeding initial data...")
			seed.SeedDatabase(config.DB)
		} else {
			log.Println("Database already seeded, skipping seeding process.")
		}
	}

	// Initialize the router
	router := mux.NewRouter()
	routes.LoadRoutes(router, config.DB)

	FRONTEND_URL := os.Getenv("FRONTEND_URL")
	if FRONTEND_URL == "" {
		log.Fatal("FRONTEND_URL environment variable is not set")
	} else {
		log.Printf("FRONTEND_URL: %s", FRONTEND_URL)
	}

	// CORS setup
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{FRONTEND_URL}, // React app origin
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	// Wrap the router with the CORS handler
	handler := c.Handler(router)

	// Create a custom HTTP server for graceful shutdown
	srv := &http.Server{
		Handler:      handler,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Run the server in a goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not start server: %v\n", err)
		}
	}()

	// Graceful shutdown setup
	waitForShutdown(srv)
}

// Graceful shutdown function
func waitForShutdown(srv *http.Server) {
	// Create a channel to listen for OS interrupt signals (e.g., SIGINT)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Block until a signal is received
	<-stop

	// Create a context with a timeout for the graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Attempt to gracefully shut down the server
	log.Println("Shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	// Close any database connections if necessary
	log.Println("Closing database connection...")
	sqlDB, err := config.DB.DB()
	if err == nil {
		sqlDB.Close()
	}

	log.Println("Server gracefully stopped.")
}
