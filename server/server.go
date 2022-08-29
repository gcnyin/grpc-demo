/*
 *
 * Copyright 2022 gcnyin.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	pb "github.com/gcnyin/grpc-demo/user"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 8848, "The server port")
)

type server struct {
	pb.UserServiceServer
}

func (s *server) GetUserById(ctx context.Context, in *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error) {
	log.Printf("Received userId: %v", in.GetUserId())
	return &pb.GetUserByIdResponse{User: &pb.User{UserId: in.GetUserId(), Username: "mock_data"}}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
