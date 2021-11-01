package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/xadrijo/vidflex-test/cmd/internal/handlers"
)

type Mapper struct {
	productHandler handlers.ProductHandler
	cartHandler handlers.CartHandler
}

func NewMapper(productHandler handlers.ProductHandler, cartHandler handlers.CartHandler) Mapper {
	return Mapper{
		productHandler: productHandler,
		cartHandler: cartHandler,
	}
}

func (m Mapper) configureMappings(router *gin.Engine) {
	router.GET("/", handlers.GetHome)
	router.GET("/product/:id", m.productHandler.GetProduct)
	router.POST("/product", m.productHandler.CreateProduct)

	router.GET("/cart/:id", m.cartHandler.GetCart)
	router.POST("/cart", m.cartHandler.CreateCart)
	router.POST("/cart/products/:id", m.cartHandler.AddProductCart)
}