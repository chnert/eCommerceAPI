package repository

import (
	"database/sql"
	"fmt"

	"eCommerceAPI/models"
)

type OrderRepo struct {
	DB *sql.DB
}

func (r *OrderRepo) CreateOrder(req *models.CheckoutRequest) error {
	// 1. Start the transaction
	// Instead of r.DB.Query, we use r.DB.Begin() to get a Transaction object (tx).
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	// 2. The Safety Net
	// If tx.commit is called at the very end, this Rollback does nothing
	// But if the function returns early because of an error, this guarantees
	// the database reverses any partial changes.
	defer tx.Rollback()

	// 3. Create the initial order record)
	var orderID int
	err = tx.QueryRow(
		"INSERT INTO orders (user_id, total_amount, status) VALUES ($1, 0, 'paid') RETURNING id",
		req.UserID,
	).Scan(&orderID)
	if err != nil {
		return err
	}

	var totalAmount float64

	// 4. Loop through the requested items
	for _, item := range req.Items {
		var price float64
		var currentStock int

		// 5. PESSIMISTIC LOCKING
		// "FOR UPDATE" locks this specific product row so no other user can buy it
		// until out transaction is finished.
		err = tx.QueryRow(
			"SELECT price, stock_quantity FROM products WHERE id = $1 FOR UPDATE",
			item.ProductID,
		).Scan(&price, &currentStock)
		if err != nil {
			return err
		}

		// 6. Business Logic: Check Stock
		if currentStock < item.Quantity {
			return fmt.Errorf("insufficient stock for product ID %d", item.ProductID)
		}

		// 7. Deduct the stock
		// We use tx.Exec() instead of QueryRow because we don't need any data returned.
		_, err = tx.Exec(
			"UPDATE products SET stock_quantity = stock_quantity - $1 WHERE id = $2",
			item.Quantity, item.ProductID,
		)
		if err != nil {
			return err
		}

		// 8. Save the receipt for this item
		_, err = tx.Exec(
			"INSERT INTO order_items (order_id, product_id, quantity, price_at_purchase) VALUES ($1, $2, $3, $4)",
			orderID, item.ProductID, item.Quantity, price,
		)
		if err != nil {
			return err
		}

		// Calculate runing total
		totalAmount += (price * float64(item.Quantity))
	}

	// 9. Update the master order with the final total
	_, err = tx.Exec(
		"UPDATE orders SET total_amount = $1 WHERE id = $2",
		totalAmount, orderID,
	)
	if err != nil {
		return err
	}

	// 10. COMMIT
	// If we reach this line, no errors occured. Save it permanently.
	return tx.Commit()
}

func (r *OrderRepo) GetOrdersByUser(userID int) ([]models.OrderResponse, error) {
	rows, err := r.DB.Query("SELECT id, total_amount, status FROM orders WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var orders []models.OrderResponse
	for rows.Next() {
		var o models.OrderResponse
		if err := rows.Scan(&o.ID, &o.TotalAmount, &o.Status); err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}
