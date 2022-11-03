package main

import (
	"context"
	"fmt"
	pb "grpcdemo/service"
	"time"

	"google.golang.org/grpc"
)

func main() {
	input := pb.InputRequest{Id: 123}
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())

	if err != nil {
		fmt.Println("dial server error: ", err)
		return
	}
	defer conn.Close()
	prodClient := pb.NewProductServiceClient(conn)
	resp, err := prodClient.GetProdStock(context.Background(), &input)
	if err != nil {
		fmt.Println("get stock error: ", err)
		return
	}
	fmt.Println(resp.Stock, resp.Name, resp.Data)

	stream, err := prodClient.PingPongStream(context.Background())
	if err != nil {
		fmt.Println("client get stream error: ", err)
		return
	}
	go handleStream(stream)
	exitChan := make(chan struct{})
	<-exitChan
}

func handleStream(stream pb.ProductService_PingPongStreamClient) {
	for {
		in := pb.InputRequest{Id: 123}
		err := stream.Send(&in)
		if err != nil {
			fmt.Println("client send stream error: ", err)
			return
		}
		resp, err := stream.Recv()
		if err != nil {
			fmt.Println("client recv stream error: ", err)
			return
		}
		fmt.Println(resp.Stock)
		time.Sleep(time.Second)
	}
}
