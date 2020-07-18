package main

import (
	"context"
	"fmt"
	"log"

	"github.com/thearyanahmed/tinkering-with-grpc/sum/sumpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatal("Error dialing up, %v", err)
	}

	defer cc.Close()

	c := sumpb.NewSumServiceClient(cc)

	callUniary(c)
}

func callUniary(c sumpb.SumServiceClient) {

	fmt.Println("Calling unary")

	req := &sumpb.SumRequest{
		Sum: &sumpb.Sum{
			A: 123,
			B: 12,
		},
	}

	res, err := c.Sum(context.Background(), req)

	if err != nil {
		log.Fatalf("Error calling sum. %v", err)
	}

	fmt.Printf("Response : %v", res.Result)
}
