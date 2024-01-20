package main

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

const (
	baseUrl = "localhost:8080"
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	//var stringTest string
	stringTest := "hello world"
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(stringTest); err != nil {
		http.Error(w, "Failed to encode data", http.StatusInternalServerError)
		return
	}
}

func main() {
	r := chi.NewRouter()

	r.Get("/", getHandler)

	err := http.ListenAndServe(baseUrl, r)
	if err != nil {
		log.Fatal(err)
	}
}
