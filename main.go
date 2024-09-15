package main

import (
	"backend_tarot/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	// Connect to the database
	// config.ConnectDatabase()

	// Automatically migrate the schema in the correct order
	// err := config.DB.AutoMigrate(
	// 	&todo.TodoCategory{}, // Migrate TodoCategory first
	// 	&todo.Todo{},         // Then migrate Todo
	// )
	// if err != nil {
	// 	log.Fatalf("Failed to migrate database: %v", err)
	// }

	// Initialize the router
	router := mux.NewRouter()
	routes.LoadRoutes(router)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // React app origin
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	// Wrap the router with the CORS handler
	handler := c.Handler(router)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", handler))
}
