package service_test

import (
	"bookcabin_project/model"
	"bookcabin_project/service"
	mock_repository "bookcabin_project/tests/mock/repository"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestVoucherService_Generate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockVoucherRepo := mock_repository.NewMockVoucherRepository(ctrl)
	voucherService := service.NewVoucherService(mockVoucherRepo)

	t.Run("successfully generate voucher", func(t *testing.T) {
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

		mockVoucherRepo.EXPECT().CreateVoucher(voucher).Return(nil)

		err := voucherService.Generate(voucher)
		assert.NoError(t, err)
	})

	t.Run("fail to generate voucher due to repository error", func(t *testing.T) {
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

		mockVoucherRepo.EXPECT().CreateVoucher(voucher).Return(errors.New("repository error"))

		err := voucherService.Generate(voucher)
		assert.Error(t, err)
		assert.Equal(t, "repository error", err.Error())
	})
}

func TestVoucherService_CheckExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockVoucherRepo := mock_repository.NewMockVoucherRepository(ctrl)
	voucherService := service.NewVoucherService(mockVoucherRepo)

	t.Run("voucher exists", func(t *testing.T) {
		flightNumber := "FL123"
		date := "2023-10-01"

		mockVoucherRepo.EXPECT().GetVoucherCountByFlighNumberAndDate(flightNumber, date).Return(int64(1), nil)

		exists, err := voucherService.CheckExist(flightNumber, date)
		assert.NoError(t, err)
		assert.True(t, exists)
	})

	t.Run("voucher does not exist", func(t *testing.T) {
		flightNumber := "FL123"
		date := "2023-10-01"

		mockVoucherRepo.EXPECT().GetVoucherCountByFlighNumberAndDate(flightNumber, date).Return(int64(0), nil)

		exists, err := voucherService.CheckExist(flightNumber, date)
		assert.NoError(t, err)
		assert.False(t, exists)
	})

	t.Run("error checking voucher existence", func(t *testing.T) {
		flightNumber := "FL123"
		date := "2023-10-01"

		mockVoucherRepo.EXPECT().GetVoucherCountByFlighNumberAndDate(flightNumber, date).Return(int64(0), errors.New("repository error"))

		exists, err := voucherService.CheckExist(flightNumber, date)
		assert.Error(t, err)
		assert.False(t, exists)
		assert.Equal(t, "repository error", err.Error())
	})
}
