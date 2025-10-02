// main.go — minimal HTTP JSON API in Go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"
)

type Response struct {
    Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    resp := Response{Message: "Hello, world from Go!"}
    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(resp); err != nil {
        http.Error(w, "failed to write response", http.StatusInternalServerError)
    }
}

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    addr := fmt.Sprintf(":%s", port)

    http.HandleFunc("/hello", helloHandler)

    log.Printf("Listening on %s", addr)
    if err := http.ListenAndServe(addr, nil); err != nil {
        log.Fatalf("Server failed: %v", err)
    }
}
