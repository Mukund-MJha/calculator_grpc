package main

import (
	"../server-grpc"
	"log"
	"net"
)

func main() {

	listen, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalf("Failed to Listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server_grpc.server{})

	log.Println("Initiating Server")
	if err = s.Serve(listen); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}
}
