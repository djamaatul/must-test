package controller

import (
	"encoding/json"
	"fmt"
	"must-test/tests/Practical/src/dto"
	"must-test/tests/Practical/src/repositories"
	"must-test/tests/Practical/src/utils"
	"net/http"
	"time"
)

type GetSummaryStruct struct {
	BookingList       []repositories.Booking     `json:"listBooking"`
	MasterConsumption []repositories.Consumption `json:"masterConsumption"`
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

	fmt.Println(periode.GoString())

	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")

	bookingListChan := make(chan []repositories.Booking)
	masterConsumptionChan := make(chan []repositories.Consumption)

	go repositories.GetBookingList(bookingListChan)
	go repositories.GetMasterConsumption(masterConsumptionChan)

	bookingList := <-bookingListChan
	masterConsumption := <-masterConsumptionChan

	var response GetSummaryStruct = GetSummaryStruct{
		BookingList:       bookingList,
		MasterConsumption: masterConsumption,
	}

	wr, _ := json.Marshal(response)

	w.Write(wr)
}
