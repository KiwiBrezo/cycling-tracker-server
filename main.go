package main

import (
	"cycling-tracker-server/Api"
	"cycling-tracker-server/Mongodb"
)

func main() {
	var endpointRouter = Api.HttpApi{}

	Mongodb.ConnectToMongoDB()

	endpointRouter.Init()

	endpointRouter.StartServer("0.0.0.0:8081")
}
