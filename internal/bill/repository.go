package bill

import (
	"context"
	"gorm.io/gorm"
)

type sqliteRepository struct {
	*gorm.DB
}

func NewBillRepository(db *gorm.DB) BillRepository {
	return &sqliteRepository{
		db,
	}
}

func (db *sqliteRepository) GetBills(ctx context.Context) ([]Bill, error) {
	var bills []Bill
	if err := db.Find(&bills).Error; err != nil {
		return nil, err
	}
	return bills, nil
}

func (db *sqliteRepository) GetBillByDay(ctx context.Context, year int, month int, day int) ([]Bill, error) {
	var bills []Bill
	if err := db.Where("year =? AND month =? AND day =?", year, month, day).Find(&bills).Error; err != nil {
		return nil, err
	}
	return bills, nil
}

func (db *sqliteRepository) GetBillByMonth(ctx context.Context, year int, month int) ([]Bill, error) {
	var bills []Bill
	if err := db.Where("year =? AND month =?", year, month).Find(&bills).Error; err != nil {
		return nil, err
	}
	return bills, nil
}

func (db *sqliteRepository) GetBillByYear(ctx context.Context, year int) ([]Bill, error) {
	var bills []Bill
	if err := db.Where("year =?", year).Find(&bills).Error; err != nil {
		return nil, err
	}
	return bills, nil
}

func (db *sqliteRepository) GetBillByID(ctx context.Context, id int) (*Bill, error) {
	var bill Bill
	if err := db.Where("id =?", id).Find(&bill).Error; err != nil {
		return nil, err
	}
	return &bill, nil
}

func (db *sqliteRepository) CreateBill(ctx context.Context, bill *Bill) error {
	return db.Create(bill).Error
}

func (db *sqliteRepository) UpdateBill(ctx context.Context, bill *Bill) error {
	return db.Save(bill).Error
}

func (db *sqliteRepository) DeleteBill(ctx context.Context, id int) error {
	return db.Where("id =?", id).Delete(&Bill{}).Error
}
