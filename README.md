# grpc-demo

1. Install `protoc`

```bash
brew install protobuf
```

2. Generate files

```bash
cd user
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ./UserService.proto
```

3. Run the server

```bash
go run server/server.go
```

4. Run the client

```bash
go run client/client.go
```
