package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"github.com/zianwar/go-rpc-calculator/calcpb"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer cc.Close()

	c := calcpb.NewCalulatorClient(cc)

	cReq := &calcpb.CalulatorRequest{
		X: 1,
		Y: 2,
	}

	cRes, err := c.Calculate(context.Background(), cReq)
	if err != nil {
		log.Printf("Server responded with error: %v\n", err)
	}

	log.Printf("Server responded with result: %v\n", cRes.GetResult())
}
