package controller

import (
	"encoding/json"
	"fmt"
	"must-test/tests/Practical/src/dto"
	"must-test/tests/Practical/src/repositories"
	"must-test/tests/Practical/src/utils"
	"net/http"
	"slices"
	"time"
)

type GetSummaryStruct struct {
	RoomName           string         `json:"roomName"`
	PercentageUsage    float64        `json:"percentageUsage"`
	NominalConsumption int            `json:"nominalConsumption"`
	TotalParticipants  int            `json:"totalParticipants"`
	Consumptions       map[string]int `json:"consumptions"`
	Currency           string         `json:"currency"`
}

func GetSummary(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body := utils.ReadBody[dto.GetSummaryDto](r)

	periode, err := time.Parse("2006-01", body.Periode)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid Payload",
		})
		return
	}

	bookingListChan := make(chan []repositories.Booking)
	masterConsumptionChan := make(chan []repositories.Consumption)

	go repositories.GetBookingList(bookingListChan)
	go repositories.GetMasterConsumption(masterConsumptionChan)

	bookingList := <-bookingListChan
	bookingList = slices.DeleteFunc(bookingList, func(e repositories.Booking) bool {
		startTime, err := time.Parse("2006-01", e.StartTime)
		if err != nil {
			return false
		}
		return !(startTime.Day() == periode.Day() && startTime.Month() == periode.Month() && startTime.Year() == periode.Year())
	})
	masterConsumption := <-masterConsumptionChan

	roomsMap := make(map[string]*GetSummaryStruct)

	for _, booking := range bookingList {
		if _, exists := roomsMap[booking.RoomName]; !exists {
			roomsMap[booking.RoomName] = &GetSummaryStruct{
				RoomName:          booking.RoomName,
				Consumptions:      make(map[string]int),
				Currency:          "Rp",
				TotalParticipants: 0,
			}
		}

		roomsMap[booking.RoomName].TotalParticipants += booking.Participants

		for _, consumption := range booking.ListConsumption {
			indexConsumption := slices.IndexFunc(masterConsumption, func(e repositories.Consumption) bool {
				return e.Name == consumption.Name
			})
			roomsMap[booking.RoomName].NominalConsumption += masterConsumption[indexConsumption].MaxPrice * booking.Participants
			roomsMap[booking.RoomName].Consumptions[consumption.Name] += booking.Participants
		}
	}

	var roomsSlice []*GetSummaryStruct = []*GetSummaryStruct{}

	for _, room := range roomsMap {
		roomsSlice = append(roomsSlice, room)
	}

	fmt.Println(periode)
	wr, _ := json.Marshal(map[string]any{
		"dataRecord": roomsSlice,
		"message":    "Sukses",
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(wr)
}
