package main

import (
	"fmt"
	"io"
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

	//callUniary(c)

	callServerStreaming(c)
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

func callServerStreaming (c greetpb.GreetServiceClient) {
	fmt.Println("Calling server streaming from client")

	req := &greetpb.GreetServerStreamRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "User",
			LastName: "Banana",
		},
	}

	responesStream, err := c.GreetServerStream(context.Background(),req)

	if err != nil {
		log.Printf("Error in response from server streamin: %v \n", err)
		return
	}

	for {
		message, err := responesStream.Recv()

		if err == io.EOF {
			fmt.Println("End of file (streaming)")
			break
		}

		if err != nil {
			log.Fatalf("error : %v ", err)
			return
		}

		fmt.Println("Message " ,message.GetResult())
	}
}