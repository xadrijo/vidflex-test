package order

import "time"

type Order struct {
	ID        int64     `json:"id"`
	CartID    int32     `json:"cart_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
