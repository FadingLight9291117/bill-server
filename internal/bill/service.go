package bill

import (
	"context"
	"github.com/samber/lo"
	"time"
)

type billService struct {
	billRepository BillRepository
}

func NewBillService(r BillRepository) BillService {
	return &billService{
		billRepository: r,
	}
}

// GetBills
// - date is {year, month, day} when return a day bills
// - date is {year, month} when return a month bills
// - date is {year} when return a year bills
// - date is {} when return all bills
func (b *billService) GetBills(ctx context.Context, date BDate) ([]interface{}, error) {
	if date.Year == 0 && (date.Month > 0 || date.Day > 0) {
		date.Year = time.Now().Year()
	} else if date.Month == 0 && date.Day > 0 {
		date.Month = int(time.Now().Month())
	}
	var bills []Bill
	var err error
	switch {
	case date.Year == 0:
		bills, err = b.billRepository.GetBills(ctx)
	case date.Month == 0:
		bills, err = b.billRepository.GetBillByYear(ctx, date.Year)
	case date.Day == 0:
		bills, err = b.billRepository.GetBillByMonth(ctx, date.Year, date.Month)
	default:
		bills, err = b.billRepository.GetBillByDay(ctx, date.Year, date.Month, date.Day)
	}

	return lo.Map(bills, func(b Bill, _ int) interface{} {
		return b.Dto()
	}), err
}

func (b *billService) GetBillByID(ctx context.Context, id int) (interface{}, error) {
	bill, err := b.billRepository.GetBillByID(ctx, id)
	if err != nil {
		return nil, err
	}
	billDto := bill.Dto()
	return &billDto, nil
}

func (b *billService) CreateBill(ctx context.Context, bill *Bill) error {
	return b.billRepository.CreateBill(ctx, bill)
}

func (b *billService) UpdateBill(ctx context.Context, bill *Bill) error {
	return b.billRepository.UpdateBill(ctx, bill)
}

func (b *billService) DeleteBill(ctx context.Context, id int) error {
	return b.billRepository.DeleteBill(ctx, id)
}
