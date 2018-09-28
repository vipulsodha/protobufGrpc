package main

import (
	"context"
	"fmt"
	"github.com/carousell/ProtobufGrpc/complex"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {

	
}

func (*server) Test(ctx context.Context, req *complexpb.ComplexMessage) (*complexpb.DummyMessage, error)  {

	return &complexpb.DummyMessage{
		Name:"Vipul Sodha",
	}, nil
}

func main()  {

	fmt.Println("Hello")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")


	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()

	complexpb.RegisterTestServiceServer(s, &server{})

	s.Serve(lis)

	fmt.Println("Running at ", lis.Addr())

}
