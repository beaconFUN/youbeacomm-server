package youbeacomm

import (
	"net/http"
)

type Beacon struct {
	ProximityUUID string `json:"proximityUUID"`
	Major         int    `json:"major"`
	Minor         int    `json:"minor"`
}

func BeaconPassedPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
