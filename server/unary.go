/*
*Unary is a simple RPC where the client sends a request to the server using the stub and waits for a response to come back, just like a normal function call. REST
 */

package main

import (
	"context"

	pb "github.com/kyleochata/mini-gerp/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello, World!"}, nil
}
