// Copyright 2020 Google LLC. All Rights Reserved.
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

package proxy

import (
	"context"
	"errors"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	pb "google.golang.org/genproto/googleapis/spanner/v1"
	"google.golang.org/grpc"
)

type spannerServer struct {
	proxy *Proxy
}

var errNotSupported = errors.New("operation not supported")

func (s *spannerServer) CreateSession(ctx context.Context, req *pb.CreateSessionRequest) (*pb.Session, error) {
	fn := s.proxy.CreateSession
	if fn == nil {
		return nil, errNotSupported
	}
	return fn(ctx, req)
}

func (s *spannerServer) BatchCreateSessions(ctx context.Context, req *pb.BatchCreateSessionsRequest) (*pb.BatchCreateSessionsResponse, error) {
	fn := s.proxy.BatchCreateSessions
	if fn == nil {
		return nil, errNotSupported
	}
	return fn(ctx, req)
}

func (s *spannerServer) GetSession(ctx context.Context, req *pb.GetSessionRequest) (*pb.Session, error) {
	fn := s.proxy.GetSession
	if fn == nil {
		return nil, errNotSupported
	}
	return fn(ctx, req)
}

func (s *spannerServer) ListSessions(ctx context.Context, req *pb.ListSessionsRequest) (*pb.ListSessionsResponse, error) {
	fn := s.proxy.ListSessions
	if fn == nil {
		return nil, errNotSupported
	}
	return fn(ctx, req)
}

func (s *spannerServer) DeleteSession(ctx context.Context, req *pb.DeleteSessionRequest) (*empty.Empty, error) {
	fn := s.proxy.DeleteSession
	if fn == nil {
		return nil, errNotSupported
	}
	return fn(ctx, req)
}

func (s *spannerServer) ExecuteSql(ctx context.Context, req *pb.ExecuteSqlRequest) (*pb.ResultSet, error) {
	fn := s.proxy.ExecuteSQL
	if fn == nil {
		return nil, errNotSupported
	}
	return fn(ctx, req)
}

func (s *spannerServer) ExecuteStreamingSql(pb *pb.ExecuteSqlRequest, ss pb.Spanner_ExecuteStreamingSqlServer) error {
	fn := s.proxy.ExecuteStreamingSQL
	if fn == nil {
		return errNotSupported
	}
	return fn(pb, ss)
}

func (s *spannerServer) ExecuteBatchDml(ctx context.Context, req *pb.ExecuteBatchDmlRequest) (*pb.ExecuteBatchDmlResponse, error) {
	fn := s.proxy.ExecuteBatchDML
	if fn == nil {
		return nil, errNotSupported
	}
	return fn(ctx, req)
}

func (s *spannerServer) Read(ctx context.Context, req *pb.ReadRequest) (*pb.ResultSet, error) {
	fn := s.proxy.Read
	if fn == nil {
		return nil, errNotSupported
	}
	return fn(ctx, req)
}

func (s *spannerServer) StreamingRead(req *pb.ReadRequest, ss pb.Spanner_StreamingReadServer) error {
	fn := s.proxy.StreamingRead
	if fn == nil {
		return errNotSupported
	}
	return fn(req, ss)
}

func (s *spannerServer) BeginTransaction(ctx context.Context, req *pb.BeginTransactionRequest) (*pb.Transaction, error) {
	fn := s.proxy.BeginTransaction
	if fn == nil {
		return nil, errNotSupported
	}
	return fn(ctx, req)
}

func (s *spannerServer) Commit(ctx context.Context, req *pb.CommitRequest) (*pb.CommitResponse, error) {
	fn := s.proxy.Commit
	if fn == nil {
		return nil, errNotSupported
	}
	return fn(ctx, req)
}

func (s *spannerServer) Rollback(ctx context.Context, req *pb.RollbackRequest) (*empty.Empty, error) {
	fn := s.proxy.Rollback
	if fn == nil {
		return nil, errNotSupported
	}
	return fn(ctx, req)
}

func (s *spannerServer) PartitionQuery(ctx context.Context, req *pb.PartitionQueryRequest) (*pb.PartitionResponse, error) {
	fn := s.proxy.PartitionQuery
	if fn == nil {
		return nil, errNotSupported
	}
	return fn(ctx, req)
}

func (s *spannerServer) PartitionRead(ctx context.Context, req *pb.PartitionReadRequest) (*pb.PartitionResponse, error) {
	fn := s.proxy.PartitionRead
	if fn == nil {
		return nil, errNotSupported
	}
	return fn(ctx, req)
}

// Proxy allows to create Google Cloud Spanner proxy servers.
// In order to override behavior, implement the functions.
// If user calls an unimplemented function, the proxy will return
// and error saying operation is not supported.
type Proxy struct {
	CreateSession       func(ctx context.Context, req *pb.CreateSessionRequest) (*pb.Session, error)
	BatchCreateSessions func(ctx context.Context, req *pb.BatchCreateSessionsRequest) (*pb.BatchCreateSessionsResponse, error)
	GetSession          func(ctx context.Context, req *pb.GetSessionRequest) (*pb.Session, error)
	ListSessions        func(ctx context.Context, req *pb.ListSessionsRequest) (*pb.ListSessionsResponse, error)
	DeleteSession       func(ctx context.Context, req *pb.DeleteSessionRequest) (*empty.Empty, error)
	ExecuteSQL          func(ctx context.Context, req *pb.ExecuteSqlRequest) (*pb.ResultSet, error)
	ExecuteStreamingSQL func(req *pb.ExecuteSqlRequest, s pb.Spanner_ExecuteStreamingSqlServer) error
	ExecuteBatchDML     func(ctx context.Context, req *pb.ExecuteBatchDmlRequest) (*pb.ExecuteBatchDmlResponse, error)
	Read                func(ctx context.Context, req *pb.ReadRequest) (*pb.ResultSet, error)
	StreamingRead       func(req *pb.ReadRequest, s pb.Spanner_StreamingReadServer) error
	BeginTransaction    func(ctx context.Context, req *pb.BeginTransactionRequest) (*pb.Transaction, error)
	Commit              func(ctx context.Context, req *pb.CommitRequest) (*pb.CommitResponse, error)
	Rollback            func(ctx context.Context, req *pb.RollbackRequest) (*empty.Empty, error)
	PartitionQuery      func(ctx context.Context, req *pb.PartitionQueryRequest) (*pb.PartitionResponse, error)
	PartitionRead       func(ctx context.Context, req *pb.PartitionReadRequest) (*pb.PartitionResponse, error)
}

// New creates a new proxy.
func New() *Proxy {
	return &Proxy{}
}

// Serve starts serving the proxy server on the given
// listener with the specified options.
func (p *Proxy) Serve(l net.Listener, opt ...grpc.ServerOption) error {
	server := grpc.NewServer(opt...)
	pb.RegisterSpannerServer(server, &spannerServer{proxy: p})
	return server.Serve(l)
}
