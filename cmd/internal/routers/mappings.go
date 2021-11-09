package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/xadrijo/vidflex-test/cmd/internal/handlers"
)

type Mapper struct {
	productHandler handlers.ProductHandler
	cartHandler handlers.CartHandler
	orderHandler handlers.OrderHandler
}

func NewMapper(
	productHandler handlers.ProductHandler,
	cartHandler handlers.CartHandler,
	orderHandler handlers.OrderHandler) Mapper {
	return Mapper{
		productHandler: productHandler,
		cartHandler: cartHandler,
		orderHandler: orderHandler,
	}
}

func (m Mapper) configureMappings(router *gin.Engine) {
	apiGroup := router.Group("/api")
	apiGroup.GET("/", handlers.GetHome)
	apiGroup.GET("/product/:id", m.productHandler.GetProduct)
	apiGroup.POST("/product", m.productHandler.CreateProduct)

	apiGroup.GET("/cart/:id", m.cartHandler.GetCart)
	apiGroup.POST("/cart", m.cartHandler.CreateCart)
	apiGroup.POST("/cart/products/:id", m.cartHandler.AddProductCart)

	apiGroup.POST("/order", m.orderHandler.CreateOrder)
	apiGroup.GET("/order/:id", m.orderHandler.GetOrder)
}