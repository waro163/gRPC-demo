package main

import (
	"context"
	"fmt"
	pb "grpcdemo/service"
	"io"
	"net"

	"google.golang.org/grpc"
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
				fmt.Println("服务接受结束")
				return err
			}
			fmt.Println("server recv error: ", err)
			return err
		}
		err = stream.Send(&pb.OutputResponse{Stock: req.Id})
		if err != nil {
			fmt.Println("server send error: ", err)
			return err
		}
	}
}

var ProductImpl = new(Product)

func main() {

	server := grpc.NewServer()
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
