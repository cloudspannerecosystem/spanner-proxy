package main

import (
	"context"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes"
	proxy "github.com/rakyll/spanner-proxy"
	pb "google.golang.org/genproto/googleapis/spanner/v1"
)

func main() {
	lis, err := net.Listen("tcp", ":9777")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	proxy := proxy.New()
	proxy.CreateSession = func(ctx context.Context, req *pb.CreateSessionRequest) (*pb.Session, error) {
		// Your own session creation...
		return &pb.Session{
			Name:       "my-first-session",
			CreateTime: ptypes.TimestampNow(),
		}, nil
	}
	proxy.Serve(lis)
}
