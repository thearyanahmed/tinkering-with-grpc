package main

import (
	"fmt"
	"log"
	"net"


	"google.golang.org/grpc"
	"github.com/thearyanahmed/tinkering-with-grpc/greet/greetpb"
)

type server struct{}

func main() {
	fmt.Println("Starting server")

	listener, err := net.Listen("tcp","0.0.0.0:50051")

	if err != nil {
		log.Fatal("Error: %v",err)
	}

	s := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(s, &server{})

	err = s.Serve(listener)

	if err != nil {
		log.Fatal("Error Serving greetpb. %v",err)
	}
}