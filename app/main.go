package main

import (
  "fmt"
  "log"
  "net/http"
  "github.com/prometheus/client_golang/prometheus/promhttp"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusOK)
  fmt.Fprintf(w, "Service is healthy and ready for traffic!")
}

func main() {
  http.HandleFunc("/health", healthHandler)
  
  http.Handle("/metrics", promhttp.Handler()) 
  
  fmt.Println("Server starting on port 8080...")
  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatalf("Server failed to start: %v", err)
  }
}