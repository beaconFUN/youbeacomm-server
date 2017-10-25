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
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if r.Header.Get("Content-Length") == "0" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var beacon PassedBeacon
	json.NewDecoder(r.Body).Decode(&beacon)

	if beacon.DeviceId == "" || beacon.Beacon.ProximityUUID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	/*
		Save proximity info to DB
	*/

	w.WriteHeader(http.StatusOK)
}
