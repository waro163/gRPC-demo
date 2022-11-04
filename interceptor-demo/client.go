package main

import (
	"context"
	"fmt"
	pb "grpcdemo/service"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	input := pb.InputRequest{Id: 123}
	conn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithChainUnaryInterceptor(AuthUnaryInterceptor))
	// conn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("dial server error: ", err)
		return
	}
	defer conn.Close()
	prodClient := pb.NewProductServiceClient(conn)

	// unary
	for i := 0; i < 3; i++ {
		resp, err := prodClient.GetProdStock(context.Background(), &input)
		if err != nil {
			fmt.Println("get stock error: ", err)
			continue
		}
		fmt.Println(resp.Stock, resp.Name, resp.Data)
		time.Sleep(time.Second)
	}

	// stream
	stream, err := prodClient.PingPongStream(context.Background())
	if err != nil {
		fmt.Println("client get stream error: ", err)
		return
	}

	exitChan := make(chan struct{})
	go handleStream(stream, exitChan)
	<-exitChan
}

func handleStream(stream pb.ProductService_PingPongStreamClient, exitCh chan struct{}) {
	count := 5
	var Num int32 = 100
	for i := 0; i < count; i++ {
		in := pb.InputRequest{Id: Num}
		err := stream.Send(&in)
		if err != nil {
			fmt.Println("client send stream error: ", err)
			return
		}
		fmt.Println("发送数据:", in.Id)
		resp, err := stream.Recv()
		if err != nil {
			fmt.Println("client recv stream error: ", err)
			return
		}
		fmt.Println("收到数据：", resp.Stock)
		Num = resp.Stock + 1
		time.Sleep(time.Second)
	}
	exitCh <- struct{}{}
}

func AuthUnaryInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	newCtx := metadata.AppendToOutgoingContext(ctx, userKey, userValue, passKey, passValue)
	return invoker(newCtx, method, req, reply, cc, opts...)
}
