package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/kyleochata/mini-gerp/proto"
)

func callSayHelloBiDiStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Bidirectional streaming started")
	stream, err := client.SayHelloBiDiStream(context.Background())
	if err != nil {
		log.Fatalf("Error while establishing stream: %v", err)
	}

	waitCh := make(chan struct{})

	// Goroutine to receive messages from the server
	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				// Stream closed by server, exit the receive loop
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving: %v", err)
			}

			log.Printf("Received: %v", message.Message)
		}
		close(waitCh)
	}()

	// Send names to the server
	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		// Send the request to the server
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending: %v", err)
		}
		time.Sleep(time.Second * 2)
	}

	// Close the send direction of the stream after sending all requests
	if err := stream.CloseSend(); err != nil {
		log.Fatalf("Error closing stream: %v", err)
	}

	// Wait for the receiving goroutine to finish
	<-waitCh
	log.Printf("Bidirectional streaming ended")
}
