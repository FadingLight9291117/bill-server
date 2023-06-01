package bill

import (
	"context"
	"gorm.io/gorm"
)

type databaseRepository struct {
	db *gorm.DB
}

func NewDatabaseRepository(db *gorm.DB) BillRepository {
	return &databaseRepository{
		db: db,
	}
}

func (db *databaseRepository) GetBills(ctx context.Context) (*[]Bill, error) {
	//TODO implement me
	panic("implement me")
}

func (db *databaseRepository) GetBillByDay(ctx context.Context, year int, month int, day int) (*[]Bill, error) {
	//TODO implement me
	panic("implement me")
}

func (db *databaseRepository) GetBillByMonth(ctx context.Context, year int, month int) (*[]Bill, error) {
	//TODO implement me
	panic("implement me")
}

func (db *databaseRepository) GetBillByYear(ctx context.Context, year int) (*[]Bill, error) {
	//TODO implement me
	panic("implement me")
}

func (db *databaseRepository) GetBillByID(ctx context.Context, id int) (*Bill, error) {
	//TODO implement me
	panic("implement me")
}

func (db *databaseRepository) CreateBill(ctx context.Context, bill *Bill) error {
	//TODO implement me
	panic("implement me")
}

func (db *databaseRepository) UpdateBill(ctx context.Context, bill *Bill) error {
	//TODO implement me
	panic("implement me")
}

func (db *databaseRepository) DeleteBill(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}
