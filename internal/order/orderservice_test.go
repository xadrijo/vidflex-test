package order

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/xadrijo/vidflex-test/internal/product"
)

func TestOrderService(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	t.Run("test get order by id", func(t *testing.T) {
		orderStorageMock := NewMockStorage(mockCtrl)
		id := int64(1)
		now := time.Now()
		products := []product.Product{
			{
				ID:        1,
				Label:     "product-1",
				Type:      "type-1",
				Url:       "",
				Weight:    54.5,
				CreatedAt: now,
				UpdatedAt: now,
			},
			{
				ID:        2,
				Label:     "product-2",
				Type:      "type-2",
				Url:       "http://someweb.com",
				Weight:    0,
				CreatedAt: now,
				UpdatedAt: now,
			},
		}
		orderStorageMock.
			EXPECT().
			GetOrderByID(id).
			Return(products, nil)

		orderService := New(orderStorageMock)
		result, err := orderService.GetOrderByID(
			context.Background(),
			id,
		)

		assert.NoError(t, err)
		assert.ElementsMatch(t, products, result)
	})

	t.Run("test insert order", func(t *testing.T) {
		orderStorageMock := NewMockStorage(mockCtrl)
		cartId := int32(1)
		now := time.Now()
		or := Order{
			CartID:    cartId,
			CreatedAt: now,
			UpdatedAt: now,
		}

		orderStorageMock.
			EXPECT().
			InsertOrder(or).
			Return(Order{
			CartID:    cartId,
			CreatedAt: now,
			UpdatedAt: now,
		}, nil)

		orderService := New(orderStorageMock)
		ord, err := orderService.InsertOrder(
			context.Background(),
			or,
			)

		assert.NoError(t, err)
		assert.Equal(t, int32(1), ord.CartID)
	})
}

// continue to complete 100% coverage in all services!!
