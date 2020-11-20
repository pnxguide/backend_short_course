package model

import (
	"log"

	"../database"
	"../form"
)

type ProductModel struct{}

func (m ProductModel) ReadAll() ([]form.Product, error) {
	db := database.Database{}
	conn, err := db.NewConnection()
	if err != nil {
		return nil, err
	}
	defer db.CloseConnection(conn)

	var products []form.Product

	rows, err := conn.Query("SELECT * FROM products")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var p form.Product
		err := rows.Scan(&p.ProductID, &p.ProductName, &p.ProductQuantity, &p.ProductPrice)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func (m ProductModel) Read(id int64) (*form.Product, error) {
	db := database.Database{}
	conn, err := db.NewConnection()
	if err != nil {
		return nil, err
	}
	defer db.CloseConnection(conn)

	rows, err := conn.Query(
		`SELECT * FROM products WHERE product_id = ?`,
		id,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var p form.Product

	if rows.Next() {
		err := rows.Scan(&p.ProductID, &p.ProductName, &p.ProductQuantity, &p.ProductPrice)
		if err != nil {
			return nil, err
		}
	}

	return &p, nil
}
