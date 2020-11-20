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
