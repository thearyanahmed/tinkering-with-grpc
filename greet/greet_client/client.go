package main

import (
	"fmt"
	"log"
	"context"

	"github.com/thearyanahmed/tinkering-with-grpc/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting client")

	cc, err := grpc.Dial("localhost:50051",grpc.WithInsecure())

	if err != nil {
		log.Fatal("Error dialing up, %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	callUniary(c)
}

func callUniary(c greetpb.GreetServiceClient) {

	fmt.Println("Calling unary")

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting {
			FirstName: "Muhaimenul Islam",
			LastName: "Ove",
		},
	}

	res, err := c.Greet(context.Background(),req)

	if err != nil {
		log.Fatalf("Error calling greet. %v",err)
	}

	fmt.Printf("Response : %v",res.Result)
}