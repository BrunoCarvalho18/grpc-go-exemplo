package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-live/pb"
	"io"
	"log"
)

func main() {

	connection, err := grpc.Dial("localhost:50055",grpc.WithInsecure())

	if err != nil{
		log.Fatalf("could not connect: %v",err)
	}

	defer connection.Close()
	client := pb.NewMathServiceClient(connection)

	request := &pb.FibonacciRequest{
		N:10,
	}

	responseStream, err := client.Fibonacci(context.Background(),request)

	for{
		stream, err := responseStream.Recv()
		 if err == io.EOF{
		 	break
		 }
		 log.Printf("Fibonacci: %v",stream.GetResult())
	}
}
