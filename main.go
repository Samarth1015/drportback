package main

import (
    "back/routes"
    "log"
    "net/http"
    "os"
)

// CORS Middleware
func corsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Set CORS headers
        w.Header().Set("Access-Control-Allow-Origin", "*") // Replace "*" with your frontend domain for production
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        
        // Handle preflight OPTIONS request
        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusOK)
            return
        }

        next.ServeHTTP(w, r)
    })
}

func main() {
    r := routes.Route() // Your router setup

    // Wrap router with CORS middleware
    handler := corsMiddleware(r)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8000" // Default port
    }

    log.Printf("Server starting on port %s...", port)
    err := http.ListenAndServe("0.0.0.0:"+port, handler)
    if err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
}
