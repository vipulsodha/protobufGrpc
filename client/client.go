package main

import (
	"context"
	"fmt"
	"github.com/carousell/ProtobufGrpc/complex"
	"google.golang.org/grpc"
	"log"
)

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := complexpb.NewTestServiceClient(conn)

	req := &complexpb.ComplexMessage {
		Name:"Vipul",
	}

	message, err := client.Test(context.Background(), req)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(message.Name)
	}

}
