package main

import (
	"context"
	"fmt"
	pb "grpcdemo/service"
	"io"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	anypb "google.golang.org/protobuf/types/known/anypb"
)

type Product struct {
	pb.UnimplementedProductServiceServer
}

func (p *Product) GetProdStock(ctx context.Context, prod *pb.InputRequest) (*pb.OutputResponse, error) {
	pid := prod.Id
	pname := prod.Name
	fmt.Println(pid, pname)
	data, _ := anypb.New(prod)
	out := &pb.OutputResponse{Stock: pid, Data: data}
	return out, nil
}

func (p *Product) PingPongStream(stream pb.ProductService_PingPongStreamServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				fmt.Println("服务接收结束")
				return err
			}
			fmt.Println("stream server recv error: ", err)
			return err
		}
		fmt.Println("收到数据：", req.Id)
		err = stream.Send(&pb.OutputResponse{Stock: req.Id + 1})
		if err != nil {
			fmt.Println("stream server send error: ", err)
			return err
		}
		fmt.Println("发送数据：", req.Id+1)
	}
}

var ProductImpl = new(Product)

func main() {

	server := grpc.NewServer(grpc.ChainUnaryInterceptor(AuthInterceptor))
	pb.RegisterProductServiceServer(server, ProductImpl)

	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println("listen error: ", err)
		return
	}

	err = server.Serve(listen)
	if err != nil {
		fmt.Println("server error: ", err)
		return
	}

}

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Not Authorization")
	}
	if md.Len() < 2 {
		return nil, status.Error(codes.Unauthenticated, "Not enough Authorization")
	}
	// fmt.Printf("%#v\n", md)
	user := md.Get(userKey)
	if len(user) < 1 {
		return nil, status.Error(codes.Unauthenticated, "Empty Authorization")
	}
	passwd := md.Get(passKey)
	if len(passwd) < 1 {
		return nil, status.Error(codes.Unauthenticated, "Wrong Authorization")
	}
	if user[0] != userValue || passwd[0] != passValue {
		return nil, status.Error(codes.Unauthenticated, "Invalid Authorization")
	}
	return handler(ctx, req)
}
