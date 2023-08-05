package Models

type RideLocation struct {
	RideId    int     `json:"rideId"`
	Timestamp uint64  `json:"timestamp"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
