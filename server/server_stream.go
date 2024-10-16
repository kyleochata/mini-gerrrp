package main

import (
	"log"
	"time"

	pb "github.com/kyleochata/mini-gerp/proto"
)

func (s *helloServer) SayHelloServerStream(req *pb.NamesList, stream pb.GreetService_SayHelloServerStreamServer) error {
	log.Printf("got request with names: %v", req.Names)
	for _, name := range req.Names {
		res := &pb.HelloResponse{Message: "Hello, " + name + "!"}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(time.Second * 2) //sim large data processing
	}
	return nil
}
