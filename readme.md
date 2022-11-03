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
server.go

run server
```bash
go run -mod=vendor ./simple/server.go
```

## client file
client.go

run client
```bash
go run -mod=vendor ./simple/client.go
```


