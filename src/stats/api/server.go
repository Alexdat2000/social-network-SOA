package api

import (
	"database/sql"
	pb "soa/stats/stats_grpc"
)

type Server struct {
	pb.UnimplementedContentServer
	Click *sql.DB
}
