package bill

import "context"

// Bill represents 'bills' object.
type Bill struct {
	ID      int     `json:"id"`
	Type    string  `json:"type"`
	Date    string  `json:"date"`
	Money   float64 `json:"money"`
	Class   string  `json:"class"`
	Label   string  `json:"label"`
	Options string  `json:"options"`
}

type BDate struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

type BillRepository interface {
	GetBills(ctx context.Context) (*[]Bill, error)
	GetBillByDate(ctx context.Context, year int, month int, day int) (*[]Bill, error)
	GetBillByMonth(ctx context.Context, year int, month int) (*[]Bill, error)
	GetBillByYear(ctx context.Context, year int) (*[]Bill, error)
	GetBillByID(ctx context.Context, id int) (*Bill, error)
	CreateBill(ctx context.Context, bill *Bill) error
	UpdateBill(ctx context.Context, bill *Bill) error
	DeleteBill(ctx context.Context, id int) error
}

type BillService interface {
	FetchBills(ctx context.Context, date BDate) (*[]Bill, error)
	FetchBillByID(ctx context.Context, id int) (*Bill, error)
	BuildBill(ctx context.Context, bill *Bill) error
	ModifyBill(ctx context.Context, id int) error
	DestroyBill(ctx context.Context, id int) error
}
