package handlers

import (
	"context"
	"github.com/xadrijo/vidflex-test/cmd/internal/product"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xadrijo/vidflex-test/cmd/internal/order"
)

type orderService interface {
	GetOrderByID(ctx context.Context, id int64) ([]product.Product, error)
	InsertOrder(ctx context.Context, c order.Order) (order.Order, error)
}

type OrderHandler struct {
	orderService orderService
}

func NewOrder(os orderService) OrderHandler {
	return OrderHandler{
		orderService: os,
	}
}

func (os OrderHandler) GetOrder(c *gin.Context) {
	orderID, err := strconv.ParseInt(c.Param("id"), 10, 32)
	ord, err := os.orderService.GetOrderByID(c, orderID)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, ord)
		return
	}

	c.IndentedJSON(http.StatusOK, ord)
}

func (os OrderHandler) CreateOrder(c *gin.Context) {
	var ord order.Order
	if err := c.ShouldBindJSON(&ord); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	log.Printf("order - handler: %d", ord.CartID)

	result, err := os.orderService.InsertOrder(c, ord)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusCreated, result)
}
