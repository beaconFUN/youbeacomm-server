package youbeacomm

import (
	"encoding/json"
	"github.com/satori/go.uuid"
	"net/http"
)

type Device struct {
	OSName    string `json:"os_name"`
	OSVersion string `json:"os_version"`
	OSLang    string `json:"os_lang"`
}

func DevicePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var body Device
	json.NewDecoder(r.Body).Decode(&body)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(uuid.NewV4().String())
}
