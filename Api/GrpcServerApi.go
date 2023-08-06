package Api

import (
	"context"
	"cycling-tracker-server/Grpc"
	"cycling-tracker-server/Services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type GrpcServerApi struct {
	Grpc.UnimplementedAddRideServiceServer
	Grpc.UnimplementedAddRideLocationServiceServer
	server      *grpc.Server
	rideService Services.RideService
}

func (api *GrpcServerApi) Init() *GrpcServerApi {
	api.server = grpc.NewServer()
	api.rideService = Services.RideService{}

	api.rideService.Init()

	Grpc.RegisterAddRideServiceServer(api.server, api)
	Grpc.RegisterAddRideLocationServiceServer(api.server, api)

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

func (api *GrpcServerApi) AddRide(context context.Context, request *Grpc.AddRideRequest) (*Grpc.AddRideResponse, error) {
	log.Println("gRPC called AddRide")

	err := api.rideService.AddRide(*request)
	if err == nil {
		return &Grpc.AddRideResponse{IsUploaded: 1}, nil
	} else {
		return &Grpc.AddRideResponse{IsUploaded: 0}, nil
	}
}

func (api *GrpcServerApi) AddRideLocation(ctx context.Context, request *Grpc.AddLocationRequest) (*Grpc.AddLocationResponse, error) {
	log.Println("gRPC called AddRideLocation")

	err := api.rideService.AddRideLocation(*request)
	if err == nil {
		return &Grpc.AddLocationResponse{IsUploaded: 1}, nil
	} else {
		return &Grpc.AddLocationResponse{IsUploaded: 0}, nil
	}
}
