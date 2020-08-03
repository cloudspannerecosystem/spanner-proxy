// Copyright 2020 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
	log.Fatal(proxy.Serve(lis))
}
