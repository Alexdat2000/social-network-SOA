package api

import (
	"soa/gateway/content_grpc"
	"soa/gateway/stats_grpc"
)

type Server struct {
	ContentAPI content_grpc.ContentClient
	StatsAPI   stats_grpc.StatsClient
}
