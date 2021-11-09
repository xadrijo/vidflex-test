//go:generate mockgen -source=orderservice.go -destination=orderservice_mocks_test.go -package=order github.com/xadrijo/vidflex-test/internal/order Storage

package order

import (
	"context"
	"log"

	"github.com/xadrijo/vidflex-test/internal/product"
)

// Storage defines the interface we expect from our db implementation
type Storage interface {
	GetOrderByID(id int64) ([]product.Product, error)
	InsertOrder(o Order) (Order, error)
}

// Service defines interaction with the DB.
type Service struct {
	Storage Storage
}

func New(storage Storage) Service {
	return Service{
		Storage: storage,
	}
}

// GetOrderByID retrieves products by order ID.
func (s Service) GetOrderByID(ctx context.Context, id int64) ([]product.Product, error) {
	products, err := s.Storage.GetOrderByID(id)
	if err != nil {
		return []product.Product{}, err
	}

	return products, nil
}

// InsertOrder inserts a new order into the DB.
func (s Service) InsertOrder(ctx context.Context, o Order) (Order, error) {
	ord, err := s.Storage.InsertOrder(o)
	if err != nil {
		log.Printf("error creating an order %s", err)
		return Order{}, err
	}

	return ord, nil
}
