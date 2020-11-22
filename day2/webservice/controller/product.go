package controller

import (
	"log"
	"net/http"
	"strconv"

	"../form"
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

func (pc ProductController) Read(c *gin.Context) {
	productModel := model.ProductModel{}

	productID, err := strconv.ParseUint(c.Query("product_id"), 10, 64)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}

	product, err := productModel.Read(productID)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, product)
}

func (pc ProductController) Add(c *gin.Context) {
	productModel := model.ProductModel{}

	var request form.ProductRequest
	err := c.BindJSON(&request)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}

	err = productModel.Add(request.ProductID,
		request.ProductName, request.ProductQuantity,
		request.ProductPrice, request.ProductTypeID)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
