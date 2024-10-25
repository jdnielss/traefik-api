package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    // Endpoint root
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        log.Println("Handling request to root endpoint")
        fmt.Fprintf(w, "Hello, World!")
    })

    // Health check endpoint
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        log.Println("Handling request to health check endpoint")
        fmt.Fprintf(w, "OK")
    })

    // API v1 endpoint
    http.HandleFunc("/api/v1", func(w http.ResponseWriter, r *http.Request) {
        log.Println("Handling request to API v1 endpoint")
        fmt.Fprintf(w, "API v1 Response")
    })

    // API v2 endpoint
    http.HandleFunc("/api/v2", func(w http.ResponseWriter, r *http.Request) {
        log.Println("Handling request to API v2 endpoint")
        fmt.Fprintf(w, "API v2 Response")
    })

    // Start the server
    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("could not start server: %s\n", err)
    }
}
