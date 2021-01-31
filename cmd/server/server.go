package main

import(
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc-live/pb"
	"grpc-live/servers"
	"log"
	"net"
)

func main() {

	grpcServer := grpc.NewServer()
	pb.RegisterMathServiceServer(grpcServer,&servers.Math{})
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp",":50055")
	if err != nil{
		log.Fatalf("Cannot start the grpc server: %v",err)
	}

	grpcServer.Serve(listener)
}

