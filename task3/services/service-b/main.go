package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/order", handleOrder)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getOrderByID(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getOrderByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	fmt.Printf("Order: %v", id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"status": true})
}
