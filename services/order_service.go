package services

import (
	"errors"
	"fmt"

	"eCommerceAPI/models"
	"eCommerceAPI/repository"
)

type OrderService struct {
	Repo *repository.OrderRepo
}

func (s *OrderService) Checkout(req *models.CheckoutRequest) error {
	maxTotalItems := 10
	maxPerItem := 3
	totalQuantity := 0
	// Rule 1: User cannot submit an empty cart.
	if len(req.Items) == 0 {
		return errors.New("cannot checkout with an empty cart")
	}

	// Rule 2: Validate each item in cart
	for _, item := range req.Items {
		// Rule 2.1: No negative or zero quantities
		if item.Quantity <= 0 {
			return fmt.Errorf("invalid quantity %d for product ID %d", item.Quantity, item.ProductID)
		}
		// Rule 2.2: Limit quantity per specific item
		if item.Quantity > maxPerItem {
			return fmt.Errorf("cannot order more than %d units of product ID %d", maxPerItem, item.ProductID)
		}
		// Keep the number of total item in the order
		totalQuantity += item.Quantity
	}

	// Rule 3: Limit total overall items in the order
	if totalQuantity > maxTotalItems {
		return fmt.Errorf("cart exceeds maximum allowed items (%d). You have %d", maxTotalItems, totalQuantity)
	}

	return s.Repo.CreateOrder(req)
}

func (s *OrderService) GetUserOrders(userID int) ([]models.OrderResponse, error) {
	if userID <= 0 {
		return nil, errors.New("invalid user ID")
	}
	return s.Repo.GetOrdersByUser(userID)
}
