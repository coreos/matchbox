package rpc

import (
	"golang.org/x/net/context"

	"github.com/mikeynap/coreos-baremetal/bootcfg/rpc/rpcpb"
	"github.com/mikeynap/coreos-baremetal/bootcfg/server"
	pb "github.com/mikeynap/coreos-baremetal/bootcfg/server/serverpb"
)

// groupServer takes a bootcfg Server and implements a gRPC GroupsServer.
type groupServer struct {
	srv server.Server
}

func newGroupServer(s server.Server) rpcpb.GroupsServer {
	return &groupServer{
		srv: s,
	}
}

func (s *groupServer) GroupPut(ctx context.Context, req *pb.GroupPutRequest) (*pb.GroupPutResponse, error) {
	_, err := s.srv.GroupPut(ctx, req)
	return &pb.GroupPutResponse{}, grpcError(err)
}

func (s *groupServer) GroupGet(ctx context.Context, req *pb.GroupGetRequest) (*pb.GroupGetResponse, error) {
	group, err := s.srv.GroupGet(ctx, req)
	return &pb.GroupGetResponse{Group: group}, grpcError(err)
}

func (s *groupServer) GroupList(ctx context.Context, req *pb.GroupListRequest) (*pb.GroupListResponse, error) {
	groups, err := s.srv.GroupList(ctx, req)
	return &pb.GroupListResponse{Groups: groups}, grpcError(err)
}
