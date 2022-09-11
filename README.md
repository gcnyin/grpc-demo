# grpc-demo

1. Install `protoc`

```bash
brew install protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

2. Generate files

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative user/user_service.proto
```

3. Run the server

```bash
go run server/server.go
```

4. Run the client

```bash
go run client/client.go
```
