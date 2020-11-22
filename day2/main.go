package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	conn, err := sql.Open("mysql",
		"dbuser:12345678@tcp(localhost:3306)/example")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	err = conn.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

type product struct {
	ProductID       uint64
	ProductName     string
	ProductQuantity uint64
	ProductPrice    float64
}

func getAllProducts(conn *sql.DB) []product {
	var products []product

	rows, err := conn.Query(`
		SELECT product_id, product_name, product_quantity,
			product_price
		FROM products;
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var p product

		err = rows.Scan(&p.ProductID, &p.ProductName,
			&p.ProductQuantity, &p.ProductPrice)
		if err != nil {
			log.Fatal(err)
		}

		products = append(products, p)
	}

	return products
}
