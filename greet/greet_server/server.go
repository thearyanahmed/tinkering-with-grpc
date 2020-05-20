package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
	"grpc-udemy-course/greet/greetpb"
)

type server struct{}

func main() {
	fmt.Println("Starting server")

	// listener, err := net.Listen("tcp","0.0.0.0:50051")

	// if err != nil {
	// 	log.Fatal("Error: %v",err)
	// }

	// s := grpc.NewServer()

	// greetpb.Re
}