package server

import (
	"../controller"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	router.GET("/ping", controller.Ping)

	productCrtl := controller.ProductController{}
	router.GET("/products", productCrtl.ReadAll)
	router.GET("/product", productCrtl.Read)

	return router
}
