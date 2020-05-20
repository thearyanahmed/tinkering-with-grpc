package main

import (
	"fmt"
	"log"
	"net"
	"context"

	"google.golang.org/grpc"
	"github.com/thearyanahmed/tinkering-with-grpc/greet/greetpb"
)

type server struct{}

func main() {
	fmt.Println("Starting server")

	listener, err := net.Listen("tcp","0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Error: %v",err)
	}

	s := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(s, &server{})

	err = s.Serve(listener)

	if err != nil {
		log.Fatalf("Error Serving greetpb. %v",err)
	}
}

func (server *server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {

	fmt.Println("greet request invoked.")
	fmt.Printf("Data: %v", req)

	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()

	message := "Hello, " + firstName + " " + lastName

	res := &greetpb.GreetResponse{
		Result: message,
	}

	return res, nil
}
