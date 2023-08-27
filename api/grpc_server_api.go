package api

import (
	"context"
	"cycling-tracker-server/grpc"
	"cycling-tracker-server/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type GrpcServerApi struct {
	grpc.UnimplementedAddRideServiceServer
	grpc.UnimplementedAddRideLocationServiceServer
	server      *grpc.Server
	rideService service.RideService
}

func (api *GrpcServerApi) Init() *GrpcServerApi {
	api.server = grpc.NewServer()
	api.rideService = service.RideService{}

	api.rideService.Init()

	grpc.RegisterAddRideServiceServer(api.server, api)
	grpc.RegisterAddRideLocationServiceServer(api.server, api)

	reflection.Register(api.server)

	return api
}

func (api *GrpcServerApi) StartServer(addr string) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("Cannot create gRPC server: ", err)
	}

	log.Printf("Start gRPC server on %s", addr)
	go func() {
		err = api.server.Serve(listener)
		if err != nil {
			log.Fatal("Cannot create gRPC server: ", err)
		}
	}()
}

func (api *GrpcServerApi) AddRide(context context.Context, request *grpc.AddRideRequest) (*grpc.AddRideResponse, error) {
	log.Println("gRPC called AddRide")

	err := api.rideService.AddRide(*request)
	if err == nil {
		return &grpc.AddRideResponse{IsUploaded: 1}, nil
	} else {
		return &grpc.AddRideResponse{IsUploaded: 0}, nil
	}
}

func (api *GrpcServerApi) AddRideLocation(ctx context.Context, request *grpc.AddLocationRequest) (*grpc.AddLocationResponse, error) {
	log.Println("gRPC called AddRideLocation")

	err := api.rideService.AddRideLocation(*request)
	if err == nil {
		return &grpc.AddLocationResponse{IsUploaded: 1}, nil
	} else {
		return &grpc.AddLocationResponse{IsUploaded: 0}, nil
	}
}
