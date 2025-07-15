package controller

import (
	"bookcabin_project/model"
	"bookcabin_project/service"
	"bookcabin_project/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VoucherController struct {
	voucherService service.VoucherService
}

func NewVoucherController(voucherService service.VoucherService) *VoucherController {
	return &VoucherController{
		voucherService: voucherService,
	}
}

type GenerateInput struct {
	ID           string `json:"id" binding:"required"`
	Name         string `json:"name" binding:"required"`
	FlightNumber string `json:"flightNumber" binding:"required"`
	Date         string `json:"date" binding:"required"`
	Aircraft     string `json:"aircraft" binding:"required"`
}

func (ac *VoucherController) Generate(c *gin.Context) {
	var input GenerateInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, exists := model.AircraftSeatMap[input.Aircraft]; !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid aircraft type"})
		return
	}

	seats := ac.getRandomizeSeat(input.Aircraft)

	voucher := model.Voucher{
		CrewName:     input.Name,
		CrewID:       input.ID,
		FlightNumber: input.FlightNumber,
		FlightDate:   input.Date,
		AircraftType: input.Aircraft,
		Seat1:        seats[0],
		Seat2:        seats[1],
		Seat3:        seats[2],
	}

	exists, err := ac.voucherService.CheckExist(voucher.FlightNumber, voucher.FlightDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "voucher already exists for this flight and date"})
		return
	}

	err = ac.voucherService.Generate(voucher)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success"})
}

type CheckInput struct {
	FlightNumber string `json:"flightNumber" binding:"required"`
	Date         string `json:"date" binding:"required"`
}

func (ac *VoucherController) Check(c *gin.Context) {
	var input CheckInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exists, err := ac.voucherService.CheckExist(input.FlightNumber, input.Date)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"exists": exists})

}

func (ac *VoucherController) getRandomizeSeat(aircraft string) []string {
	seats := model.AircraftSeatMap[aircraft]
	seats = utils.RandomizeSlice(seats)
	return seats[:3]
}
