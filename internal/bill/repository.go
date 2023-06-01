package bill

import (
	"context"
	"gorm.io/gorm"
)

type sqliteRepository struct {
	db *gorm.DB
}

func NewDatabaseRepository(db *gorm.DB) BillRepository {
	return &sqliteRepository{
		db: db,
	}
}

func (db *sqliteRepository) GetBills(ctx context.Context) (*[]Bill, error) {
	//TODO implement me
	panic("implement me")
}

func (db *sqliteRepository) GetBillByDay(ctx context.Context, year int, month int, day int) (*[]Bill, error) {
	//TODO implement me
	panic("implement me")
}

func (db *sqliteRepository) GetBillByMonth(ctx context.Context, year int, month int) (*[]Bill, error) {
	//TODO implement me
	panic("implement me")
}

func (db *sqliteRepository) GetBillByYear(ctx context.Context, year int) (*[]Bill, error) {
	//TODO implement me
	panic("implement me")
}

func (db *sqliteRepository) GetBillByID(ctx context.Context, id int) (*Bill, error) {
	//TODO implement me
	panic("implement me")
}

func (db *sqliteRepository) CreateBill(ctx context.Context, bill *Bill) error {
	//TODO implement me
	panic("implement me")
}

func (db *sqliteRepository) UpdateBill(ctx context.Context, bill *Bill) error {
	//TODO implement me
	panic("implement me")
}

func (db *sqliteRepository) DeleteBill(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}
