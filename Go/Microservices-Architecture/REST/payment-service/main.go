// payment-service/main.go

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Payment struct {
	OrderID string  `json:"order_id"`
	Amount  float64 `json:"amount"`
	Status  string  `json:"status"`
}

func processPayment(orderID string, amount float64) Payment {
	// Simulate payment processing logic
	return Payment{
		OrderID: orderID,
		Amount:  amount,
		Status:  "Success",
	}
}

func paymentHandler(w http.ResponseWriter, r *http.Request) {
	var orderRequest struct {
		OrderID string  `json:"order_id"`
		Amount  float64 `json:"amount"`
	}

	err := json.NewDecoder(r.Body).Decode(&orderRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	payment := processPayment(orderRequest.OrderID, orderRequest.Amount)

	// Communicate with the Order Service to update order status
	updateOrderStatus(payment.OrderID, "Paid")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payment)
}

func updateOrderStatus(orderID, status string) {
	// Implement communication logic with the Order Service to update order status
	fmt.Printf("Updating order %s status to %s\n", orderID, status)
}

func main() {
	http.HandleFunc("/payment", paymentHandler)

	port := 8081
	fmt.Printf("Payment Service listening on :%d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
