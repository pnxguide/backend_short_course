package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	ProductID       int
	ProductName     string
	ProductQuantity int
	ProductPrice    float64
}

func main() {
	db, err := sql.Open("mysql",
		"maria_dba:dbapassword@tcp(127.0.0.1:3306)/mydb")
	if err != nil {
		log.Fatal("Cannot connect to the database")
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot ping to the database")
	}

	log.Println(GetAllProducts(db))
}

func GetAllProducts(db *sql.DB) []Product {
	var products []Product

	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ProductID, &p.ProductName, &p.ProductQuantity, &p.ProductPrice)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, p)
	}

	return products
}
