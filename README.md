# Client Stream

A simple client streaming app to send stream of data every 10th second

### Generate proto

`protoc --proto_path=proto proto/books.proto --go_out=server --go-grpc_out=server`

`protoc --proto_path=proto proto/books.proto --go_out=client --go-grpc_out=client`

### Run server

`cd server && go run main.go`

### Run client

`cd client && go run main.go`
