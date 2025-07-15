package repository

import (
	"bookcabin_project/model"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateVoucher(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}
	defer db.Close()

	mock.ExpectQuery("select sqlite_version()").
		WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("3.36.0"))

	gormDB, err := gorm.Open(sqlite.Dialector{Conn: db}, &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open gorm DB: %v", err)
	}

	repo := NewVoucherRepository(gormDB)

	voucher := model.Voucher{
		CrewName:     "Sarah",
		CrewID:       "456",
		FlightNumber: "FL123",
		FlightDate:   "2023-10-01",
		AircraftType: "ATR",
		Seat1:        "1A",
		Seat2:        "1B",
		Seat3:        "1C",
	}

	mock.ExpectBegin()
	mock.ExpectQuery("INSERT INTO `vouchers` \\(`crew_name`,`crew_id`,`flight_number`,`flight_date`,`aircraft_type`,`seat1`,`seat2`,`seat3`,`created_at`\\) VALUES \\(\\?,\\?,\\?,\\?,\\?,\\?,\\?,\\?,\\?\\) RETURNING `id`").
		WithArgs(voucher.CrewName, voucher.CrewID, voucher.FlightNumber, voucher.FlightDate, voucher.AircraftType, voucher.Seat1, voucher.Seat2, voucher.Seat3, sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	err = repo.CreateVoucher(voucher)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet mock expectations: %v", err)
	}
}

func TestCreateVoucher_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}
	defer db.Close()

	mock.ExpectQuery("select sqlite_version()").
		WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("3.36.0"))

	gormDB, err := gorm.Open(sqlite.Dialector{Conn: db}, &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open gorm DB: %v", err)
	}

	repo := NewVoucherRepository(gormDB)

	voucher := model.Voucher{
		CrewName:     "Sarah",
		CrewID:       "456",
		FlightNumber: "FL123",
		FlightDate:   "2023-10-01",
		AircraftType: "ATR",
		Seat1:        "1A",
		Seat2:        "1B",
		Seat3:        "1C",
	}

	mock.ExpectBegin()
	mock.ExpectQuery("INSERT INTO `vouchers` \\(`crew_name`,`crew_id`,`flight_number`,`flight_date`,`aircraft_type`,`seat1`,`seat2`,`seat3`,`created_at`\\) VALUES \\(\\?,\\?,\\?,\\?,\\?,\\?,\\?,\\?,\\?\\) RETURNING `id`").
		WithArgs(voucher.CrewName, voucher.CrewID, voucher.FlightNumber, voucher.FlightDate, voucher.AircraftType, voucher.Seat1, voucher.Seat2, voucher.Seat3, sqlmock.AnyArg()).
		WillReturnError(fmt.Errorf("db insert error"))
	mock.ExpectRollback()

	err = repo.CreateVoucher(voucher)
	if err == nil {
		t.Errorf("expected error but got nil")
	} else if err.Error() != "db insert error" {
		t.Errorf("unexpected error: got %v, want %v", err, "db insert error")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet mock expectations: %v", err)
	}
}

func TestGetVoucherCountByFlighNumberAndDate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}
	defer db.Close()

	mock.ExpectQuery("select sqlite_version()").
		WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("3.36.0"))

	gormDB, err := gorm.Open(sqlite.Dialector{Conn: db}, &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open gorm DB: %v", err)
	}

	repo := NewVoucherRepository(gormDB)

	flightNumber := "FL123"
	date := "2023-10-01"

	mock.ExpectQuery("SELECT count\\(\\*\\) FROM `vouchers` WHERE flight_number = \\? AND flight_date = \\?").
		WithArgs(flightNumber, date).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(5))

	count, err := repo.GetVoucherCountByFlighNumberAndDate(flightNumber, date)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if count != 5 {
		t.Errorf("unexpected count: got %d, want %d", count, 5)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet mock expectations: %v", err)
	}
}

func TestGetVoucherCountByFlighNumberAndDate_NoResults(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}
	defer db.Close()

	mock.ExpectQuery("select sqlite_version()").
		WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("3.36.0"))

	gormDB, err := gorm.Open(sqlite.Dialector{Conn: db}, &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open gorm DB: %v", err)
	}

	repo := NewVoucherRepository(gormDB)

	flightNumber := "FL123"
	date := "2023-10-01"

	mock.ExpectQuery("SELECT count\\(\\*\\) FROM `vouchers` WHERE flight_number = \\? AND flight_date = \\?").
		WithArgs(flightNumber, date).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))

	count, err := repo.GetVoucherCountByFlighNumberAndDate(flightNumber, date)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if count != 0 {
		t.Errorf("unexpected count: got %d, want %d", count, 0)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet mock expectations: %v", err)
	}
}

func TestGetVoucherCountByFlighNumberAndDate_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}
	defer db.Close()

	mock.ExpectQuery("select sqlite_version()").
		WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("3.36.0"))

	gormDB, err := gorm.Open(sqlite.Dialector{Conn: db}, &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open gorm DB: %v", err)
	}

	repo := NewVoucherRepository(gormDB)

	flightNumber := "FL123"
	date := "2023-10-01"

	mock.ExpectQuery("SELECT count\\(\\*\\) FROM `vouchers` WHERE flight_number = \\? AND flight_date = \\?").
		WithArgs(flightNumber, date).
		WillReturnError(fmt.Errorf("db select error"))

	_, err = repo.GetVoucherCountByFlighNumberAndDate(flightNumber, date)
	if err == nil {
		t.Errorf("expected error but got nil")
	} else if err.Error() != "db select error" {
		t.Errorf("unexpected error: got %v, want %v", err, "db select error")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet mock expectations: %v", err)
	}
}
