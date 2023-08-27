package models

type Ride struct {
	RideId    int    `json:"rideId"`
	UserId    int    `json:"userId"`
	TimeStart uint64 `json:"timeStart"`
	TimeEnd   uint64 `json:"timeEnd"`
	Duration  uint64 `json:"duration"`
}
