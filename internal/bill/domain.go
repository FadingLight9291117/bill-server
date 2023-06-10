package bill

import (
	"context"
	"gorm.io/gorm"
)

// Bill represents 'bills' object.
type Bill struct {
	gorm.Model
	Year    int
	Month   int
	Day     int
	Money   float64
	Label   string
	Options string
	Type    string
}

type BillDTO struct {
	ID      int     `json:"id"`
	Year    int     `json:"year"`
	Month   int     `json:"month"`
	Day     int     `json:"day"`
	Money   float64 `json:"money"`
	Label   string  `json:"label"`
	Options string  `json:"options"`
	Type    string  `json:"type"`
}

func (b *Bill) Dto() BillDTO {
	return BillDTO{
		ID:      int(b.ID),
		Year:    b.Year,
		Month:   b.Month,
		Day:     b.Day,
		Money:   b.Money,
		Label:   b.Label,
		Options: b.Options,
		Type:    b.Type,
	}
}

// BDate 用于判断获取bill的时间范围
type BDate struct {
	Year  int `validate:"min=0"`
	Month int `validate:"min=0,max=12"`
	Day   int `validate:"min=0,max=31"`
}

type BillRepository interface {
	GetBills(ctx context.Context) ([]Bill, error)
	GetBillByDay(ctx context.Context, year int, month int, day int) ([]Bill, error)
	GetBillByMonth(ctx context.Context, year int, month int) ([]Bill, error)
	GetBillByYear(ctx context.Context, year int) ([]Bill, error)
	GetBillByID(ctx context.Context, id int) (*Bill, error)
	CreateBill(ctx context.Context, bill *Bill) error
	UpdateBill(ctx context.Context, bill *Bill) error
	DeleteBill(ctx context.Context, id int) error
}

type BillService interface {
	GetBills(ctx context.Context, date BDate) ([]interface{}, error)
	GetBillByID(ctx context.Context, id int) (interface{}, error)
	CreateBill(ctx context.Context, bill *Bill) error
	UpdateBill(ctx context.Context, bill *Bill) error
	DeleteBill(ctx context.Context, id int) error
}
