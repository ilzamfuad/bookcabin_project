package model

import (
	"fmt"
)

var (
	AircraftSeatMap = map[string][]string{
		"ATR":            generateSeatMap(18, []string{"A", "C", "D", "F"}),
		"Airbus 320":     generateSeatMap(32, []string{"A", "B", "C", "D", "E", "F"}),
		"Boeing 737 Max": generateSeatMap(32, []string{"A", "B", "C", "D", "E", "F"}),
	}
)

type Voucher struct {
	ID           int    `gorm:"primaryKey;autoIncrement" json:"id"`
	CrewName     string `gorm:"type:text;not null;" json:"crew_name"`
	CrewID       string `gorm:"type:text;not null;" json:"crew_id"`
	FlightNumber string `gorm:"type:text;not null;" json:"flight_number"`
	FlightDate   string `gorm:"type:text;not null;" json:"flight_date"`
	AircraftType string `gorm:"type:text;not null;" json:"aircraft_type"`
	Seat1        string `gorm:"type:text;not null;" json:"seat1"`
	Seat2        string `gorm:"type:text;not null;" json:"seat2"`
	Seat3        string `gorm:"type:text;not null;" json:"seat3"`
	CreatedAt    string `gorm:"type:datetime;default:CURRENT_TIMESTAMP" json:"created_at"`
}

func generateSeatMap(maxNumber int, word []string) []string {
	var seats []string
	for i := 1; i <= maxNumber; i++ {
		for _, seat := range word {
			seatCode := fmt.Sprintf("%d%s", i, seat)
			seats = append(seats, seatCode)
		}
	}
	return seats
}
