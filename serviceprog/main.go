package main

import (
	"log"
	"net/http"

	"./config"
	"./controller"
)

func init() {
	config.Init()
}

func main() {
	http.HandleFunc("/ping", controller.Ping)

	productCtrl := controller.ProductController{}
	http.HandleFunc("/products", productCtrl.ReadAll)

	log.Fatal(http.ListenAndServe(":4000", nil))
}
