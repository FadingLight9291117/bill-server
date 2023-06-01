package bill

import (
	"context"
	"gorm.io/gorm"
)

// Bill represents 'bills' object.
type Bill struct {
	gorm.Model
	Type    string
	Date    string
	Money   float64
	Class   string
	Label   string
	Options string
}

// BDate 用于判断获取bill的时间范围
type BDate struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

type BillRepository interface {
	GetBills(ctx context.Context) (*[]Bill, error)
	GetBillByDay(ctx context.Context, year int, month int, day int) (*[]Bill, error)
	GetBillByMonth(ctx context.Context, year int, month int) (*[]Bill, error)
	GetBillByYear(ctx context.Context, year int) (*[]Bill, error)
	GetBillByID(ctx context.Context, id int) (*Bill, error)
	CreateBill(ctx context.Context, bill *Bill) error
	UpdateBill(ctx context.Context, bill *Bill) error
	DeleteBill(ctx context.Context, id int) error
}

type BillService interface {
	GetBills(ctx context.Context, date BDate) (*[]Bill, error)
	GetBillByID(ctx context.Context, id int) (*Bill, error)
	CreateBill(ctx context.Context, bill *Bill) error
	UpdateBill(ctx context.Context, bill *Bill) error
	DeleteBill(ctx context.Context, id int) error
}
