package models

// CheckoutRequest is the JSON the client sends us
type CheckoutRequest struct {
	UserID int            `json:"user_id"`
	Items  []CheckoutItem `json:"items"`
}

type CheckoutItem struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type OrderResponse struct {
	ID          int     `json:"id"`
	TotalAmount float64 `json: "total_amount"`
	Status      string  `json: "status"`
}
