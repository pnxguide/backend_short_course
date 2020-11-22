package controller

import (
	"log"
	"net/http"

	"../model"
	"github.com/gin-gonic/gin"
)

type ProductController struct{}

func (pc ProductController) ReadAll(c *gin.Context) {
	productModel := model.ProductModel{}
	products, err := productModel.ReadAll()
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, products)
}
