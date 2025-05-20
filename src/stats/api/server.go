package api

import (
	pb "soa/stats/stats_grpc"
)

type Server struct {
	pb.UnimplementedContentServer
}
