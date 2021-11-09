package db

import (
	"database/sql/driver"
	"errors"
	"testing"
	"time"

	"github.com/xadrijo/vidflex-test/internal/order"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
)

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func TestOrder_Insert(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	s, _ := New(db)

	tests := []struct {
		name    string
		s       Storage
		order   order.Order
		mock    func()
		want    int32
		wantErr bool
	}{
		{
			//When everything works as expected
			name: "OK",
			s:    s,
			order: order.Order{
				CartID: int32(1),
			},
			mock: func() {
				rows := sqlxmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("INSERT INTO customer_order").WithArgs(1, AnyTime{}, AnyTime{}).WillReturnRows(rows)
			},
			want: int32(1),
		},
		{
			name:  "Empty Fields",
			s:     s,
			order: order.Order{
				CartID: int32(1),
			},
			mock: func() {
				_ = sqlxmock.NewRows([]string{"id"})
				mock.ExpectQuery("INSERT INTO users").
					WithArgs(1, AnyTime{}, AnyTime{}).
					WillReturnError(errors.New("error inserting into the db"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.s.InsertOrder(tt.order)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error new = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && got.CartID != tt.want {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
// En real app we should continue creating test until to reach 100% of coverage
