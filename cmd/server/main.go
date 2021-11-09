package main

import (
	"github.com/xadrijo/vidflex-test/cmd/internal/cart"
	"github.com/xadrijo/vidflex-test/cmd/internal/order"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/xadrijo/vidflex-test/cmd/internal/db"
	"github.com/xadrijo/vidflex-test/cmd/internal/handlers"
	"github.com/xadrijo/vidflex-test/cmd/internal/product"
	"github.com/xadrijo/vidflex-test/cmd/internal/routers"
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
	productHandler := handlers.New(productService)
	cartService := cart.New(storageProvider)
	cartHandler := handlers.NewCart(cartService)
	orderService := order.New(storageProvider)
	orderHandler := handlers.NewOrder(orderService)
	mapper := routers.NewMapper(
		productHandler,
		cartHandler,
		orderHandler,
		)
	router = routers.CreateRouter(mapper)

	return nil
}
