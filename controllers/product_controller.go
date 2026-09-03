package controllers

import (
	"encoding/json"
	"net/http"

	"eCommerceAPI/services"
	"eCommerceAPI/models"
)

type ProductController struct {
	Service *services.ProductService
}

func (c *ProductController) GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	// Call the service
	products, err := c.Service.GetAllProducts()
	if err != nil {
		http.Error(w, "Failed to get products", http.StatusInternalServerError)
		return
	}

	// Send JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (c *ProductController) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var p models.Product

	// 1. Decode the JSON body into the  struct
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		return
	}

	// 2. Call the service to save it
	if err := c.Service.CreateProduct(&p); err != nil {
		http.Error(w, "Failed to create product", http.StatusInternalServerError)
		return
	}

	// 3. Return the created product (which now includes the DB-generated ID) with a 201 created status
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}
