// package main

// import (
// 	"context"
// 	"log"
// 	"time"

// 	pb "github.com/kyleochata/mini-gerp/proto"
// )

// func callSayHello(client pb.GreetServiceClient) {
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	defer cancel()

//		res, err := client.SayHello(ctx, &pb.NoParam{})
//		if err != nil {
//			log.Fatalf("Failed to call SayHello: %v", err)
//		}
//		log.Printf("Response from server: %s", res.Message)
//	}
package main
