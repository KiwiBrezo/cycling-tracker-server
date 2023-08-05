package main

import (
	"cycling-tracker-server/Api"
	"cycling-tracker-server/Mongodb"
)

// protoc --proto_path=Grpc/proto --go_out=. --go-grpc_out=. AddRide.proto
// protoc --proto_path=Grpc/proto --go_out=. --go-grpc_out=. AddRideLocation.proto
func main() {
	var endpointRouter = Api.HttpApi{}
	var grpcApi = Api.GrpcServerApi{}

	Mongodb.ConnectToMongoDB()

	endpointRouter.Init()
	grpcApi.Init()

	grpcApi.StartServer("0.0.0.0:8081")
	endpointRouter.StartServer("0.0.0.0:8080")
}
