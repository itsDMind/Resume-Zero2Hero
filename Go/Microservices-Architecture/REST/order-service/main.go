// order-service/main.go

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Order struct {
	ID     string  `json:"id"`
	Amount float64 `json:"amount"`
	Status string  `json:"status"`
}

var orders = make(map[string]Order)

func createOrderHandler(w http.ResponseWriter, r *http.Request) {
	var order Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order.ID = generateID()
	order.Status = "Pending"
	orders[order.ID] = order

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

func getOrderHandler(w http.ResponseWriter, r *http.Request) {
	orderID := r.URL.Query().Get("id")
	order, found := orders[orderID]
	if !found {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

func generateID() string {
	// Implement your own ID generation logic
	return "12345"
}

func main() {
	http.HandleFunc("/orders", createOrderHandler)
	http.HandleFunc("/order", getOrderHandler)

	port := 8080
	fmt.Printf("Order Service listening on :%d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
