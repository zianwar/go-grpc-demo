package main

import (
	"context"
	"log"
	"net"

	"github.com/zianwar/go-rpc-calculator/calcpb"
	"google.golang.org/grpc"
)

type calculatorServer struct{}

func (cs *calculatorServer) Calculate(ctx context.Context, cReq *calcpb.CalulatorRequest) (*calcpb.CalulatorResponse, error) {
	x := cReq.GetX()
	y := cReq.GetY()

	log.Printf("Received request to calculate %d + %d\n", x, y)
	cRes := &calcpb.CalulatorResponse{
		Result: x + y,
	}

	return cRes, nil
}

func main() {
	addr := "localhost:50051"

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}

	cServer := &calculatorServer{}

	s := grpc.NewServer()
	calcpb.RegisterCalulatorServer(s, cServer)

	log.Printf("Server listeing at %s\n", addr)
	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
