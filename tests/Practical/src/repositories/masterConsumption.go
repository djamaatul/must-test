package repositories

type Consumption struct {
	CreatedAt string `json:"createdAt"`
	Name      string `json:"name"`
	MaxPrice  int    `json:"maxPrice"`
	Id        int    `json:"id"`
}

func GetMasterConsumption(response chan []Consumption) {
	res := Requester[[]Consumption]("https://6686cb5583c983911b03a7f3.mockapi.io/api/dummy-data/masterJenisKonsumsi")

	response <- res
}
