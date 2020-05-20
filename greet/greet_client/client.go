package main

import (
	"fmt"
	"log"

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

	fmt.Printf("Created client, %f", c)


}