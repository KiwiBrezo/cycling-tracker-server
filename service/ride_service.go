package service

import (
	"cycling-tracker-server/grpc"
	"cycling-tracker-server/models"
	"cycling-tracker-server/repositorie"
	"log"
)

type RideService struct {
	rideRepository repositorie.RideRepository
}

func (r *RideService) Init() {
	r.rideRepository = repositorie.RideRepository{}
}

func (r RideService) AddRide(request grpc.AddRideRequest) (err error) {
	var newRide = models.Ride{
		RideId:    int(request.RideId),
		UserId:    int(request.UserId),
		TimeStart: request.GetTimeStart(),
		TimeEnd:   request.GetTimeStop(),
		Duration:  request.Duration,
	}

	_, err = r.rideRepository.SaveRideToDatabase(newRide)
	if err != nil {
		log.Printf("There was an error inserting the new ride: %v", err)
		return
	}

	return nil
}

func (r RideService) AddRideLocation(request grpc.AddLocationRequest) (err error) {
	var newRideLocation = models.RideLocation{
		RideId:    int(request.RideId),
		Timestamp: request.Timestamp,
		Latitude:  request.Latitude,
		Longitude: request.Longitude,
	}

	_, err = r.rideRepository.SaveRideLocationToDatabase(newRideLocation)
	if err != nil {
		log.Printf("There was an error inserting the new ride location: %v", err)
		return
	}

	return nil
}
