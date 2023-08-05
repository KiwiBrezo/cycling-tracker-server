package Services

import (
	"cycling-tracker-server/Grpc"
	"cycling-tracker-server/Models"
	"cycling-tracker-server/Repositories"
	"log"
)

type RideService struct {
	rideRepository Repositories.RideRepository
}

func (r *RideService) Init() {
	r.rideRepository = Repositories.RideRepository{}
}

func (r RideService) AddRide(request Grpc.AddRideRequest) (err error) {
	var newRide = Models.Ride{
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

func (r RideService) AddRideLocation(request Grpc.AddLocationRequest) (err error) {
	var newRideLocation = Models.RideLocation{
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
