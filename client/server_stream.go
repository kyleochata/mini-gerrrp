package main

// import (
// 	"context"
// 	"io"
// 	"log"

// 	pb "github.com/kyleochata/mini-gerp/proto"
// )

// func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {
// 	log.Printf("Streaming started")
// 	stream, err := client.SayHelloServerStream(context.Background(), names)
// 	if err != nil {
// 		log.Fatalf("Failed to call SayHelloServerStream: %v", err)
// 	}
// 	for {
// 		res, err := stream.Recv()
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			log.Fatalf("Failed to receive response: %v", err)
// 		}
// 		log.Printf("Response from server: %s", res.Message)

// 	}
// 	log.Printf("Streaming finished")
// }
