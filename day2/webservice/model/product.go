package model

import (
	"../database"
	"../form"
)

type ProductModel struct{}

func (pm ProductModel) ReadAll() ([]form.Product, error) {
	conn, err := database.NewConnection()
	if err != nil {
		return nil, err
	}
	defer database.CloseConnection(conn)

	var products []form.Product

	rows, err := conn.Query(`
		SELECT product_id, product_name, product_quantity,
			product_price
		FROM products;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p form.Product

		err = rows.Scan(&p.ProductID, &p.ProductName,
			&p.ProductQuantity, &p.ProductPrice)
		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}
