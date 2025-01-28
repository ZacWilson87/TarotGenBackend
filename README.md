# Tarot Card App Backend API

Welcome to the **Tarot Card App Backend API**, a evolving system that powers a unique tarot card reading application. This API is built with **Go** to ensure performance, scalability, and ease of integration with the frontend.

## About the Project

The **Tarot Card App** offers users a way to perform tarot card readings and interact with customizable decks. Eventually, it will feature the ability to generate tarot card art dynamically, with specific focus on accurately representing cards like the "6 of Swords" or "9 of Cups," where precise item counts are essential.

This backend API serves as the foundation of the application, managing users, tarot readings, card and deck data, and preparing for advanced art generation capabilities.

---

## Features

- **User Management**: WIP
  - Authentication and authorization using secure methods like Auth0.
  - User data stored with a PostgreSQL database.

- **Tarot Readings**: 
  - Perform detailed readings with multiple spread types.
  - Associate readings with specific users for history tracking.

- **Deck and Card Management**: WIP
  - Create, manage, and customize tarot decks.
  - Validate generated cards to ensure they match predefined tarot standards.

- **Art Generation (In Development)**: WIP
  - Utilize AI models to dynamically create tarot card art.
  - Fine-tuning planned to improve control over item placement, such as exact counts of swords, cups, or pentagrams.

---

## Technology Stack

- **Language**: Go (Golang)
- **Database**: PostgreSQL (with GORM for ORM)
- **Authentication**: Auth0
- **API Routing**: Gorilla Mux

---

## Roadmap

1. **Current Development**:
   - Implement tarot card readings and spreads.
   - Establish user authentication and referral system.
   - Build endpoints for deck and card management.

2. **Future Development**:
   - Integrate more precise AI-powered art generation for tarot cards.
   - Fine-tune AI models to achieve high accuracy in visual elements like item counts. i.e. "7 of Swords" or "10 of Cups"

3. **Long-Term Goals**:
   - Enable a dropshipping service for users to print and ship their custom decks.
   - Develop a self-service platform for users to create and order their own tarot card designs.

---

## How to Run the API

### Prerequisites

1. Install **Go** (latest version recommended).
2. Set up **PostgreSQL** and configure the database connection in the project.
3. Clone this repository.
4. Set up .env file based off .env-example
