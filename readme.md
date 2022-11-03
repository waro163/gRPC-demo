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

## project explain
in simple-demo/ dir, it is a simple demo;

in auth-demo/ dir, it is a demo using authorization;

## server file
server.go

run server
```bash
go run -mod=vendor ./xxxx/server.go
```

## client file
client.go

run client
```bash
go run -mod=vendor ./xxxx/client.go
```


