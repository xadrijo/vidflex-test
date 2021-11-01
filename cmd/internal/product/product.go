package product

import (
	"time"
)

// Product should contain the definitions of a product
type Product struct {
	ID        int32
	Label     string
	Type      string
	Url       string
	Weight    float32
	CreatedAt time.Time
	UpdatedAt time.Time
}
