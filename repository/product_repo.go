package repository

import (
	"database/sql"
	"eCommerceAPI/models"
)

type ProductRepo struct {
	DB *sql.DB
}

func (r *ProductRepo) GetAll() ([]models.Product, error) {
	rows, err := r.DB.Query("SELECT id, name, price, stock_quantity FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.StockQuantity); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (r *ProductRepo) Create(p *models.Product) error {
	sqlStatement := `
		INSERT INTO products (name, price, stock_quantity)
		VALUES ($1, $2, $3)
		RETURNING id;`

	// We use QueryRow because we expect exactly one ID back
	err := r.DB.QueryRow(sqlStatement, p.Name, p.Price, p.StockQuantity).Scan(&p.ID)
	return err
}
