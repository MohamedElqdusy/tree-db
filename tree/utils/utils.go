package utils

import (
	"encoding/json"
	"net/http"
)

func WriteError(w *http.ResponseWriter, err error) {
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")
	(*w).WriteHeader(400)
	err = json.NewEncoder(*w).Encode(err)
}
