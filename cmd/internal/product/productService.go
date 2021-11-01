package product

import (
	"context"
	"log"
)

// Storage defines the interface we expect from our db implementation
type Storage interface {
	GetProductByID(id int32) (Product, error)
	InsertProduct(p Product) (Product, error)
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

// GetProductByID retrieves a product based on the ID.
func (s Service) GetProductByID(ctx context.Context, id int32) (Product, error) {
	p, err := s.Storage.GetProductByID(id)
	if err != nil {
		return Product{}, err
	}

	return p, nil
}

// InsertProduct inserts a new rocket into the DB.
func (s Service) InsertProduct(ctx context.Context, p Product) (Product, error) {
	p, err := s.Storage.InsertProduct(p)
	if err != nil {
		log.Printf("error creating product %s", err)
		return Product{}, err
	}

	return p, nil
}

