package bill

import (
	"context"
	"database/sql"
)

const (
	QUERY_GET_BILLS          = "SELECT * FROM Bills"
	QUERY_GET_BILLS_BY_DAY   = ""
	QUERY_GET_BILLS_BY_MONTH = ""
	QUERY_GET_BILLS_BY_YEAR  = ""
	QUERY_GET_BILL_BY_ID     = ""
)

type mariaDBRepository struct {
	mariadb *sql.DB
}

func NewMariaDBRepository(mariadb *sql.DB) BillRepository {
	return &mariaDBRepository{
		mariadb: mariadb,
	}
}

func (m *mariaDBRepository) GetBills(ctx context.Context) (*[]Bill, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mariaDBRepository) GetBillByDay(ctx context.Context, year int, month int, day int) (*[]Bill, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mariaDBRepository) GetBillByMonth(ctx context.Context, year int, month int) (*[]Bill, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mariaDBRepository) GetBillByYear(ctx context.Context, year int) (*[]Bill, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mariaDBRepository) GetBillByID(ctx context.Context, id int) (*Bill, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mariaDBRepository) CreateBill(ctx context.Context, bill *Bill) error {
	//TODO implement me
	panic("implement me")
}

func (m *mariaDBRepository) UpdateBill(ctx context.Context, bill *Bill) error {
	//TODO implement me
	panic("implement me")
}

func (m *mariaDBRepository) DeleteBill(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}
