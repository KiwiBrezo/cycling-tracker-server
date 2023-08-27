package repositorie

import (
	"context"
	"cycling-tracker-server/models"
	"cycling-tracker-server/mongo_db"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type RideRepository struct {
}

func (r *RideRepository) SaveRideToDatabase(ride models.Ride) (insertedId string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := mongo_db.GetCollection("rides").InsertOne(ctx, ride)
	if err != nil {
		return
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r *RideRepository) SaveRideLocationToDatabase(rideLocation models.RideLocation) (insertedId string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := mongo_db.GetCollection("ride_locations").InsertOne(ctx, rideLocation)
	if err != nil {
		return
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}
