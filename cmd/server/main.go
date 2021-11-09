package main

import (
	"github.com/xadrijo/vidflex-test/internal/cart"
	"github.com/xadrijo/vidflex-test/internal/db"
	handlers2 "github.com/xadrijo/vidflex-test/internal/handlers"
	"github.com/xadrijo/vidflex-test/internal/order"
	"github.com/xadrijo/vidflex-test/internal/product"
	routers2 "github.com/xadrijo/vidflex-test/internal/routers"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func Run() error {
	log.Println("Starting up Vidflex-APi Service")

	err := manageDependencies()
	if err != nil {
		return err
	}

	_ = router.Run(":8000")
	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}

func manageDependencies() error {
	storageProvider, err := db.New()
	if err != nil {
		return err
	}

	err = storageProvider.Migrate()
	if err != nil {
		log.Println("Failed to run migrations")
		return err
	}

	productService := product.New(storageProvider)
	productHandler := handlers2.New(productService)
	cartService := cart.New(storageProvider)
	cartHandler := handlers2.NewCart(cartService)
	orderService := order.New(storageProvider)
	orderHandler := handlers2.NewOrder(orderService)
	mapper := routers2.NewMapper(
		productHandler,
		cartHandler,
		orderHandler,
		)
	router = routers2.CreateRouter(mapper)

	return nil
}
