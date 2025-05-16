package api

import "soa/gateway/content_grpc"

type Server struct {
	ContentAPI content_grpc.ContentClient
}
