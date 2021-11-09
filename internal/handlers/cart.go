package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xadrijo/vidflex-test/internal/cart"
	"github.com/xadrijo/vidflex-test/internal/product"
)

type cartService interface {
	GetCartByID(ctx context.Context, id int32) ([]product.Product, error)
	InsertCart(ctx context.Context, c cart.Cart) (cart.Cart, error)
	AddProductInCart(ctx context.Context, c cart.CartProduct) (cart.CartProduct, error)
}

type CartHandler struct {
	cartService cartService
}

func NewCart(cs cartService) CartHandler {
	return CartHandler{
		cartService: cs,
	}
}

func (cs CartHandler) GetCart(c *gin.Context) {
	cartID, err := strconv.ParseInt(c.Param("id"), 10, 32)
	log.Println(cartID)
	products, err := cs.cartService.GetCartByID(c, int32(cartID))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, products)
		return
	}

	c.IndentedJSON(http.StatusOK, products)
}

func (cs CartHandler) CreateCart(c *gin.Context) {
	var ct cart.Cart
	if err := c.ShouldBindJSON(&ct); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	result, err := cs.cartService.InsertCart(c, ct)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusCreated, result)
}

func (cs CartHandler) AddProductCart(c *gin.Context) {
	productID, err := strconv.ParseInt(c.Param("id"), 10, 32)

	var cp cart.CartProduct
	if err := c.ShouldBindJSON(&cp); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	cp.ProductID = int32(productID)

	result, err := cs.cartService.AddProductInCart(c, cp)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusCreated, result)
}
