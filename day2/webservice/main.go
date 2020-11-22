package main

import (
	"net/http"

	"./config"
	"./controller"
	"./database"
)

func main() {
	config.Init()
	database.Init()

	http.HandleFunc("/ping", controller.Ping)
	http.HandleFunc("/products", controller.GetAllProducts)

	http.ListenAndServe("localhost:4000", nil)
}
