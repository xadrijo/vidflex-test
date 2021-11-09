package cart

import "time"

// Cart should contain the definitions of a shopping cart
type Cart struct {
	ID        int32
	CreatedAt time.Time
	UpdatedAt time.Time
}

// CartProduct should contain the definitions of a cart product struct
type CartProduct struct {
	ID        int32     `json:"id"`
	CartID    int32     `json:"cart_id"`
	ProductID int32     `json:"product_id"`
	Quantity  int32     `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
