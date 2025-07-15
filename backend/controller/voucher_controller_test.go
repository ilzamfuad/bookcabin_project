package controller_test

import (
	"bookcabin_project/controller"
	mock_service "bookcabin_project/tests/mock/service"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestVoucherController_Generate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_service.NewMockVoucherService(ctrl)
	voucherController := controller.NewVoucherController(mockService)

	router := gin.Default()
	router.POST("/api/generate", voucherController.Generate)

	t.Run("success", func(t *testing.T) {
		input := controller.GenerateInput{
			ID:           "123",
			Name:         "John Doe",
			FlightNumber: "FL123",
			Date:         "2023-10-01",
			Aircraft:     "ATR",
		}
		bodyJSON, _ := json.Marshal(input)

		mockService.EXPECT().CheckExist(input.FlightNumber, input.Date).Return(false, nil)
		mockService.EXPECT().Generate(gomock.Any()).Return(nil)

		req, _ := http.NewRequest(http.MethodPost, "/api/generate", bytes.NewBuffer(bodyJSON))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("bad request - invalid JSON", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/api/generate", bytes.NewBuffer([]byte("invalid json")))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("bad request - invalid aircraft type", func(t *testing.T) {
		input := controller.GenerateInput{
			ID:           "123",
			Name:         "John Doe",
			FlightNumber: "FL123",
			Date:         "2023-10-01",
			Aircraft:     "InvalidAircraft",
		}
		bodyJSON, _ := json.Marshal(input)

		req, _ := http.NewRequest(http.MethodPost, "/api/generate", bytes.NewBuffer(bodyJSON))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("voucher already exists", func(t *testing.T) {
		input := controller.GenerateInput{
			ID:           "123",
			Name:         "John Doe",
			FlightNumber: "FL123",
			Date:         "2023-10-01",
			Aircraft:     "ATR",
		}
		bodyJSON, _ := json.Marshal(input)

		mockService.EXPECT().CheckExist(input.FlightNumber, input.Date).Return(true, nil)

		req, _ := http.NewRequest(http.MethodPost, "/api/generate", bytes.NewBuffer(bodyJSON))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("internal service  error  - check error", func(t *testing.T) {
		input := controller.GenerateInput{
			ID:           "123",
			Name:         "John Doe",
			FlightNumber: "FL123",
			Date:         "2023-10-01",
			Aircraft:     "ATR",
		}
		bodyJSON, _ := json.Marshal(input)

		mockService.EXPECT().CheckExist(input.FlightNumber, input.Date).Return(false, errors.New("service error"))

		req, _ := http.NewRequest(http.MethodPost, "/api/generate", bytes.NewBuffer(bodyJSON))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("internal server error - service error", func(t *testing.T) {
		input := controller.GenerateInput{
			ID:           "123",
			Name:         "John Doe",
			FlightNumber: "FL123",
			Date:         "2023-10-01",
			Aircraft:     "ATR",
		}
		bodyJSON, _ := json.Marshal(input)

		mockService.EXPECT().CheckExist(input.FlightNumber, input.Date).Return(false, nil)
		mockService.EXPECT().Generate(gomock.Any()).Return(errors.New("service error"))

		req, _ := http.NewRequest(http.MethodPost, "/api/generate", bytes.NewBuffer(bodyJSON))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestVoucherController_Check(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_service.NewMockVoucherService(ctrl)
	voucherController := controller.NewVoucherController(mockService)

	router := gin.Default()
	router.POST("/api/check", voucherController.Check)

	t.Run("success - voucher exists", func(t *testing.T) {
		input := controller.CheckInput{
			FlightNumber: "FL123",
			Date:         "2023-10-01",
		}
		bodyJSON, _ := json.Marshal(input)

		mockService.EXPECT().CheckExist(input.FlightNumber, input.Date).Return(true, nil)

		req, _ := http.NewRequest(http.MethodPost, "/api/check", bytes.NewBuffer(bodyJSON))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("success - voucher does not exist", func(t *testing.T) {
		input := controller.CheckInput{
			FlightNumber: "FL123",
			Date:         "2023-10-01",
		}
		bodyJSON, _ := json.Marshal(input)

		mockService.EXPECT().CheckExist(input.FlightNumber, input.Date).Return(false, nil)

		req, _ := http.NewRequest(http.MethodPost, "/api/check", bytes.NewBuffer(bodyJSON))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("bad request - invalid JSON", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/api/check", bytes.NewBuffer([]byte("invalid json")))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("internal server error - service error", func(t *testing.T) {
		input := controller.CheckInput{
			FlightNumber: "FL123",
			Date:         "2023-10-01",
		}
		bodyJSON, _ := json.Marshal(input)

		mockService.EXPECT().CheckExist(input.FlightNumber, input.Date).Return(false, errors.New("service error"))

		req, _ := http.NewRequest(http.MethodPost, "/api/check", bytes.NewBuffer(bodyJSON))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
