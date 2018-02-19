package youbeacomm

import (
	"encoding/json"
	"fmt"
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

var beaconLogs []Beacon = make([]Beacon, 0, 3)

func getBeaconLog(s int) Beacon {
	if s < 0 || len(beaconLogs) <= s {
		return Beacon{}
	}

	return beaconLogs[s]
}

func GetLatestBeaconLog(s int) Beacon {
	return getBeaconLog(len(beaconLogs) - s - 1)
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
	beaconLogs = append(beaconLogs, beacon.Beacon)
	fmt.Println(beaconLogs)
	w.WriteHeader(http.StatusOK)
}
