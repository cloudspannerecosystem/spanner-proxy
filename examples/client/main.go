package main

import (
	"context"
	"fmt"
	"log"

	pb "google.golang.org/genproto/googleapis/spanner/v1"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	conn, err := grpc.Dial("localhost:9777", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewSpannerClient(conn)
	session, err := client.CreateSession(ctx, &pb.CreateSessionRequest{
		Database: "first-db",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(session)
}
