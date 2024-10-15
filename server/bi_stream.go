package main

import (
	"io"
	"log"

	pb "github.com/kyleochata/mini-gerp/proto"
)

func (s *helloServer) SayHelloBiDiStream(stream pb.GreetService_SayHelloBiDiStreamServer) error {
	log.Println("Bidirectional stream started on server")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("Stream closed by client, returning...")
			return nil
		}
		if err != nil {
			log.Printf("Error receiving from stream: %v", err)
			return err
		}

		// Log the received request
		log.Printf("Got request with name: %v", req.Name)

		// Prepare and send a stream back to client
		res := &pb.HelloResponse{Message: "Hello, " + req.Name + "!"}
		if err := stream.Send(res); err != nil {
			log.Printf("Error sending to stream: %v", err)
			return err
		}
		log.Printf("Response sent to client: %v", res.Message)
	}
}
