package service

import (
	"bookcabin_project/model"
	"bookcabin_project/repository"
)

type VoucherService interface {
	Generate(voucher model.Voucher) error
	CheckExist(flight_number, date string) (bool, error)
}

type voucherService struct {
	voucherRepo repository.VoucherRepository
}

func NewVoucherService(voucherRepo repository.VoucherRepository) VoucherService {
	return &voucherService{
		voucherRepo: voucherRepo,
	}
}

func (us *voucherService) Generate(voucher model.Voucher) error {
	return us.voucherRepo.CreateVoucher(voucher)
}

func (us *voucherService) CheckExist(flight_number, date string) (bool, error) {
	count, err := us.voucherRepo.GetVoucherCountByFlighNumberAndDate(flight_number, date)
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}
