package Repositories

import (
	"context"
	"cycling-tracker-server/Models"
	"cycling-tracker-server/Mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type RideRepository struct {
}

func (r *RideRepository) SaveRideToDatabase(ride Models.Ride) (insertedId string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := Mongodb.GetCollection("rides").InsertOne(ctx, ride)
	if err != nil {
		return
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r *RideRepository) SaveRideLocationToDatabase(rideLocation Models.RideLocation) (insertedId string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := Mongodb.GetCollection("ride_locations").InsertOne(ctx, rideLocation)
	if err != nil {
		return
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}
