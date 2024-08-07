package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ReadBody[V any](r *http.Request) V {

	body, error := io.ReadAll(r.Body)

	if error != nil {
		panic(error)
	}

	var parsedBody V

	json.Unmarshal(body, &parsedBody)

	return parsedBody
}
