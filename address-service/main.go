package main

import (
	"context"
	"fmt"
	pb "github.com/dwiprastyoisworo/go-grpc-simple/proto/address"
	"google.golang.org/grpc"
	"log"
	"net"
)

type addressServer struct {
	pb.UnimplementedAddressServiceServer
}

func (s *addressServer) GetAddressByUserID(ctx context.Context, req *pb.AddressRequest) (*pb.AddressResponse, error) {
	// Contoh data statis
	addresses := map[string]*pb.AddressResponse{
		"123": {
			UserId:  "123",
			Street:  "Jl. Sudirman No.123",
			City:    "Jakarta",
			ZipCode: "10230",
		},
	}

	if addr, ok := addresses[req.UserId]; ok {
		return addr, nil
	}
	return nil, fmt.Errorf("address not found")
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterAddressServiceServer(s, &addressServer{})

	fmt.Println("Address Service running on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
