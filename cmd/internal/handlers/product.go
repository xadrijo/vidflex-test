package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xadrijo/vidflex-test/cmd/internal/product"
)

type productService interface {
	GetProductByID(ctx context.Context, id int32) (product.Product, error)
	InsertProduct(ctx context.Context, p product.Product) (product.Product, error)
}

type ProductHandler struct {
	productService productService
}

func New(ps productService) ProductHandler {
	return ProductHandler{
		productService: ps,
	}
}

func (p ProductHandler) GetProduct(c *gin.Context) {
	productID, err := strconv.ParseInt(c.Param("id"), 10, 32)
	prod, err := p.productService.GetProductByID(c, int32(productID))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, prod)
		return
	}

	c.IndentedJSON(http.StatusOK, prod)
}

func (p ProductHandler) CreateProduct(c *gin.Context) {
	var prod product.Product
	if err := c.ShouldBindJSON(&prod); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	pr, err := p.productService.InsertProduct(c, prod)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusCreated, pr)
}
