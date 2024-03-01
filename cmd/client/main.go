package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "go-grpc-stream/chatgrpc"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewChatServiceClient(conn)

	stream, err := c.Chat(context.Background())
	if err != nil {
		log.Fatalf("could not chat: %v", err)
	}

	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a message : %v", err)
			}
			log.Printf("Received: %s", in.Message)
		}
	}()

	for i := 0; i < 5; i++ {
		if err := stream.Send(&pb.ChatMessage{User: "User1", Message: "Hello from the client!"}); err != nil {
			log.Fatalf("Failed to send a message: %v", err)
		}
		time.Sleep(1 * time.Second)
	}
	stream.CloseSend()
}
