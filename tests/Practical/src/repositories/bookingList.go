package repositories

import (
	"encoding/json"
	"net/http"
)

type ConsumptionList struct {
	Name string `json:"name"`
}

type Booking struct {
	BookingDate     string            `json:"bookingDate"`
	OfficeName      string            `json:"officeName"`
	StartTime       string            `json:"startTime"`
	EndTime         string            `json:"endTime"`
	ListConsumption []ConsumptionList `json:"listConsumption"`
	Participants    int               `json:"participants"`
	RoomName        string            `json:"roomName"`
	Id              int               `json:"id"`
}

func Requester[T any](url string) T {
	var data T
	res, err := http.Get(url)

	if err != nil {
		panic("Internal Server Error")
	}

	json.NewDecoder(res.Body).Decode(&data)

	return data
}

func GetBookingList(response chan []Booking) {

	res := Requester[[]Booking]("https://66876cc30bc7155dc017a662.mockapi.io/api/dummy-data/bookingList")

	response <- res
}
