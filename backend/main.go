package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/trends", trendsHandler)

	port := "8080"
	fmt.Printf("🚀 Server running on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

func trendsHandler(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("../data/processed/mocked.json")
	if err != nil {
		http.Error(w, "Failed to read trends file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Contetn-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
