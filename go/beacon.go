package youbeacomm

import (
	"encoding/json"
	"net/http"
)

type Beacon struct {
	ProximityUUID string `json:"proximityUUID"`
	Major         int    `json:"major"`
	Minor         int    `json:"minor"`
}

type PassedBeacon struct {
	DeviceId string `json:"device"`
	Beacon   Beacon `json:"beacon"`
}

func BeaconPassedPost(w http.ResponseWriter, r *http.Request) {

	if r.Header.Get("Content-Length") == "0" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var beacon PassedBeacon
	json.NewDecoder(r.Body).Decode(&beacon)

	/*
		Save proximity info to DB
	*/

	w.WriteHeader(http.StatusOK)
}
