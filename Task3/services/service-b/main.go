package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/calculate_order", handleCalculateOrder)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleCalculateOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getCalculateOrdereByID(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getCalculateOrdereByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	response, err := http.Get("http://service-a:8080/order/" + id)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response.Status, string(body))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"status": true})
}
