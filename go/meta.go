package youbeacomm

import (
	"encoding/json"
	"net/http"
)

type Meta struct {
	APIVersions []string `json:"available_api_versions"`
}

var api_versions = []string{
	"1.0.0-alpha",
	"1.0.0-alpha.2017092905",
	"1.0.0-alpha.2017103001",
}

func RootGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(Meta{api_versions})
}
