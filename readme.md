## grpc demo

## init

```bash
go mod tidy
go mod vendor
```

## proto file

/pb_file/product.proto

## compile proto file

```bash
cd grpc_demo
protoc --go_out=. --go-grpc_out=. ./pb_file/product.proto
# this will generate pb.go and grpc.pb.go file in ./service
```

## server file

pb_server.go

run server
```bash
go run -mod=vendor pb_server.go
```

## client file

copy ./service/ all file to ./pb_client/
./pb_client/client.go

run client
```bash
go run -mod=vendor ./pb_client/client.go
```


