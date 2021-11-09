//go:generate mockgen -destination=cartservice_mocks_test.go -package=cart github.com/xadrijo/

package cart

import (
	"context"
	"log"

	"github.com/xadrijo/vidflex-test/internal/product"
)

// Storage defines the interface we expect from our db implementation
type Storage interface {
	GetCartByID(id int32) ([]product.Product, error)
	InsertCart(c Cart) (Cart, error)
	AddProductInCart(cp CartProduct) (CartProduct, error)
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

// GetCartByID retrieves a cart based on the ID.
func (s Service) GetCartByID(ctx context.Context, id int32) ([]product.Product, error) {
	ps, err := s.Storage.GetCartByID(id)
	if err != nil {
		return []product.Product{}, err
	}

	return ps, nil
}

// InsertCart inserts a new cart into the DB.
func (s Service) InsertCart(ctx context.Context, c Cart) (Cart, error) {
	c, err := s.Storage.InsertCart(c)
	if err != nil {
		log.Printf("error creating cart %s", err)
		return Cart{}, err
	}

	return c, nil
}

// AddProductInCart inserts a new product in a specific cart.
func (s Service) AddProductInCart(ctx context.Context, cp CartProduct) (CartProduct, error) {
	c, err := s.Storage.AddProductInCart(cp)
	if err != nil {
		log.Printf("error inserting product in cart %s", err)
		return CartProduct{}, err
	}

	return c, nil
}
