package repository

import (
	"bookcabin_project/model"

	"gorm.io/gorm"
)

type VoucherRepository interface {
	CreateVoucher(voucher model.Voucher) error
	GetVoucherCountByFlighNumberAndDate(flight_number string, date string) (int64, error)
}

type voucherRepository struct {
	db *gorm.DB
}

func NewVoucherRepository(db *gorm.DB) VoucherRepository {
	return &voucherRepository{
		db: db,
	}
}

func (ur *voucherRepository) CreateVoucher(voucher model.Voucher) error {
	err := ur.db.Create(&voucher).Error
	if err != nil {
		return err
	}
	return nil
}

func (ur *voucherRepository) GetVoucherCountByFlighNumberAndDate(flight_number string, date string) (int64, error) {
	var count int64
	err := ur.db.Model(model.Voucher{}).Where("flight_number = ? AND flight_date = ?", flight_number, date).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}
