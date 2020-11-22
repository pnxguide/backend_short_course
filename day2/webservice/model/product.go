package model

import (
	"errors"

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

func (pm ProductModel) Read(id uint64) (form.Product, error) {
	conn, err := database.NewConnection()
	if err != nil {
		return form.Product{}, err
	}
	defer database.CloseConnection(conn)

	var product form.Product

	rows, err := conn.Query(`
		SELECT product_id, product_name, product_quantity,
			product_price
		FROM products
		WHERE product_id = ?;
	`, id)
	if err != nil {
		return form.Product{}, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&product.ProductID, &product.ProductName,
			&product.ProductQuantity, &product.ProductPrice)
		if err != nil {
			return form.Product{}, err
		}

		return product, nil
	}

	return form.Product{}, errors.New("no products belong to the id")
}
