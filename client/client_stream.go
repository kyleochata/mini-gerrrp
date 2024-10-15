package main

// import (
// 	"context"
// 	"log"
// 	"time"

// 	pb "github.com/kyleochata/mini-gerp/proto"
// )

// func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList) {
// 	log.Printf("Client streaming started")
// 	stream, err := client.SayHelloClientStream(context.Background())
// 	if err != nil {
// 		log.Fatalf("couldn't send names list: %v", err)
// 	}
// 	for _, name := range names.Names {
// 		req := &pb.HelloRequest{
// 			Name: name,
// 		}
// 		if err := stream.Send(req); err != nil {
// 			log.Fatalf("error while sending %v", err)
// 		}
// 		log.Printf("Sent the request with name: %s", name)
// 		time.Sleep(time.Second * 2)
// 	}
// 	res, err := stream.CloseAndRecv()
// 	log.Printf("Client streaming ended")
// 	if err != nil {
// 		log.Fatalf("couldn't receive response: %v", err)
// 	}
// 	log.Printf("%v", res.Messages)
// }
