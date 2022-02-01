package main

import "go-micro.dev/v4"

func main() {

	// create a new service
	service := micro.NewService(
		micro.Name("helloworld2"),
	)

	// initialise flags
	service.Init()

	// start the service
	service.Run()
}
