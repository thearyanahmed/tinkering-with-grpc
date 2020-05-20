package main

import (
	"fmt"
	"log"
	"net"
	"context"

	"google.golang.org/grpc"
	"github.com/thearyanahmed/tinkering-with-grpc/sum/sumpb"
)

type server struct{}

func main() {
	fmt.Println("Starting server")

	listener, err := net.Listen("tcp","0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Error: %v",err)
	}

	s := grpc.NewServer()

	sumpb.RegisterSumServiceServer(s, &server{})

	err = s.Serve(listener)

	if err != nil {
		log.Fatalf("Error Serving sumpb. %v",err)
	}
}

func (server *server) Sum(ctx context.Context, req *sumpb.SumRequest) (*sumpb.SumResponse, error) {

	fmt.Println("sum request invoked.")
	fmt.Printf("Data: %v", req)

	a := req.GetSum().GetA()
	b := req.GetSum().GetB()

	sum := a + b

	res := &sumpb.SumResponse{
		Result: sum,
	}

	return res, nil
}
