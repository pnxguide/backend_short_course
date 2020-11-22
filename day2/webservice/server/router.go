package server

import (
	"../controller"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	router.GET("/ping", controller.Ping)

	productController := controller.ProductController{}
	router.GET("/products", productController.ReadAll)

	return router
}
