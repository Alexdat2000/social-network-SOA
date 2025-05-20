package api

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "soa/stats/stats_grpc"
)

func (s *Server) Stats(context.Context, *pb.PostStatsRequest) (*pb.PostStats, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stats not implemented")
}
func (s *Server) Daily(context.Context, *pb.DailyRequest) (*pb.DailyStats, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Daily not implemented")
}
func (s *Server) TopPosts(context.Context, *pb.TopRequest) (*pb.TopPostList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TopPosts not implemented")
}
func (s *Server) TopUsers(context.Context, *pb.TopRequest) (*pb.TopUserList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TopUsers not implemented")
}
