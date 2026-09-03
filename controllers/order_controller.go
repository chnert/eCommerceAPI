package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"eCommerceAPI/models"
	"eCommerceAPI/services"
)

type OrderController struct {
	Service *services.OrderService
}

func (c *OrderController) CheckoutHandler(w http.ResponseWriter, r *http.Request) {
	var req models.CheckoutRequest

	// Parse the JSON body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON format.", http.StatusBadRequest)
	}

	// Call the service layer
	err := c.Service.Checkout(&req)
	// Handle the business logic errors
	if err != nil {
		// Set the header to JSON so the client knows how to read the error
		w.Header().Set("Content-Type", "application/json")
		// Send the HTTP 400 Bad Request status code
		w.WriteHeader(http.StatusBadRequest)

		// Wrap the error string in a JSON object and send it
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	// Success Resonse
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // HTTP 201
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Order placed successfully.",
	})
}

func (c *OrderController) GetOrdersHandler(w http.ResponseWriter, r *http.Request) {
	// Extract ?user_id= from the URL
	userIDstr := r.URL.Query().Get("user_id")
	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Call the service
	orders, err := c.Service.GetUserOrders(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Return the JSON array
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(orders)
}
