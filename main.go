package main

import (
    "back/routes"
    "log"
    "net/http"
    "github.com/gorilla/handlers"  // Import gorilla handlers package
    "os"
)

func main() {
    route := routes.Route() // Your router setup

    // Setup CORS using gorilla handlers
    corsHandler := handlers.CORS(
        handlers.AllowedOrigins([]string{"*"}),  // Allows all origins for development
        handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}), // Allowed HTTP methods
        handlers.AllowedHeaders([]string{"Content-Type", "Authorization", "X-Requested-With"}), // Allowed headers
    )

    // Start the server with CORS enabled
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Default port
    }

    log.Printf("Server starting on port %s...", port)
    err := http.ListenAndServe(":"+port, corsHandler(route))  // Wrap your route with CORS handler
    if err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
}
