package main

import (
	"cycling-tracker-server/api"
	"cycling-tracker-server/mongo_db"
)

// protoc --proto_path=grpc/proto --go_out=. --go-grpc_out=. add_ride.proto
// protoc --proto_path=grpc/proto --go_out=. --go-grpc_out=. add_ride_location.proto
func main() {
	var endpointRouter = api.HttpApi{}
	var grpcApi = api.GrpcServerApi{}

	mongo_db.ConnectToMongoDB()

	endpointRouter.Init()
	grpcApi.Init()

	grpcApi.StartServer("0.0.0.0:8081")
	endpointRouter.StartServer("0.0.0.0:8080")
}
