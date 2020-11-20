package controller

import (
	"log"
	"net/http"
	"strconv"

	"../model"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	m *model.ProductModel
}

func (pc ProductController) Model() *model.ProductModel {
	// singleton
	if pc.m == nil {
		pc.m = &model.ProductModel{}
	}
	return pc.m
}

func (pc ProductController) ReadAll(c *gin.Context) {
	products, err := pc.Model().ReadAll()
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, products)
}

func (pc ProductController) Read(c *gin.Context) {
	idParam, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	product, err := pc.Model().Read(idParam)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, product)
}
