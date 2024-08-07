package server

import (
	controller "must-test/tests/Practical/src/controllers"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/summary", controller.GetSummary)

	err := http.ListenAndServe("localhost:5000", nil)
	if err != nil {
		panic("Listen Server Failed")
	}
}
