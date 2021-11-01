package db

import (
	"errors"
	"fmt"
	"github.com/xadrijo/vidflex-test/cmd/internal/cart"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/xadrijo/vidflex-test/cmd/internal/product"
)

type Storage struct {
	db *sqlx.DB
}

// New - return a new store or error
func New() (Storage, error) {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSL_MODE")

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		dbHost,
		dbPort,
		dbUsername,
		dbName,
		dbPassword,
		dbSSLMode,
	)

	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return Storage{}, err
	}
	return Storage{
		db: db,
	}, nil
}

// GetProductByID - retrieves a product from the database by id
func (s Storage) GetProductByID(id int32) (product.Product, error) {
	var p product.Product

	query := "SELECT id, label, type, url, weight, createdAt, updatedAt FROM product WHERE id=$1;"
	row := s.db.QueryRow(
		query,
		id,
	)
	err := row.Scan(&p.ID, &p.Label, &p.Type, &p.Url, &p.Weight, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		log.Print(err.Error())
		return product.Product{}, err
	}
	return p, nil
}

func (s Storage) InsertProduct(prod product.Product) (product.Product, error) {
	now := time.Now()
	query := "INSERT INTO product (label, type, url, weight, createdAt, updatedAt) " +
		"VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;"

	lastInsertId := 0
	err := s.db.QueryRow(
		query,
		prod.Label, prod.Type, prod.Url, prod.Weight, now, now,
	).Scan(&lastInsertId)

	if err != nil {
		log.Print(err.Error())
		return product.Product{}, errors.New("failed to insert into database")
	}

	return product.Product{
		ID:        int32(lastInsertId),
		Label:     prod.Label,
		Type:      prod.Type,
		Url:       prod.Url,
		Weight:    prod.Weight,
		CreatedAt: prod.CreatedAt,
		UpdatedAt: prod.UpdatedAt,
	}, nil
}

func (s Storage) GetCartByID(id int32) ([]product.Product, error) {
	query := "SELECT cp.id, p.id, p.label, p.type, p.url, p.weight, p.createdAt, p.updatedAt " +
		"FROM cart_product AS cp " +
		"INNER JOIN product AS p " +
		"ON p.id = cp.product_id " +
		"WHERE cp.cart_id = $1;"

	rows, err := s.db.Query(query, id)
	if err != nil {
		return []product.Product{}, errors.New("something wrong")
	}
	defer rows.Close()

	var products []product.Product
	for rows.Next() {
		var cartID int32
		var productID int32
		var label string
		var productType string
		var url string
		var weight float32
		var createdAt time.Time
		var updatedAt time.Time

		var p product.Product
		err := rows.Scan(&cartID, &productID, &label, &productType, &url,
			&weight, &createdAt, &updatedAt)
		if err != nil {
			log.Print(err.Error())
			return products, err
		}

		p.ID = productID
		p.Label = label
		p.Type = productType
		p.Url = url
		p.Weight = weight
		p.CreatedAt = createdAt
		p.UpdatedAt = updatedAt

		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		return products, err
	}

	return products, nil
}

func (s Storage) InsertCart(c cart.Cart) (cart.Cart, error) {
	now := time.Now()
	query := "INSERT INTO cart (createdAt, updatedAt) " +
		"VALUES ($1, $2) RETURNING id;"

	lastInsertId := 0
	err := s.db.QueryRow(
		query,
		now, now,
	).Scan(&lastInsertId)

	if err != nil {
		log.Print(err.Error())
		return cart.Cart{}, errors.New("failed to insert into database")
	}

	return cart.Cart{
		ID:        int32(lastInsertId),
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}, nil
}

func (s Storage) AddProductInCart(cp cart.CartProduct) (cart.CartProduct, error) {
	now := time.Now()
	query := "INSERT INTO cart_product (cart_id, product_id, quantity, createdAt, updatedAt) " +
		"VALUES ($1, $2, $3, $4, $5) RETURNING id;"

	lastInsertId := 0
	err := s.db.QueryRow(
		query,
		cp.CartID, cp.ProductID, cp.Quantity, now, now,
	).Scan(&lastInsertId)

	if err != nil {
		log.Print(err.Error())
		return cart.CartProduct{}, errors.New("failed to insert into database")
	}

	return cart.CartProduct{
		ID:        int32(lastInsertId),
		CartID:    cp.CartID,
		ProductID: cp.ProductID,
		Quantity:  cp.Quantity,
		CreatedAt: cp.CreatedAt,
		UpdatedAt: cp.UpdatedAt,
	}, nil
}
