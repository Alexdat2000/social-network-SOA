package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"soa/stats/api"
	pb "soa/stats/stats_grpc"
)

func main() {
	port := flag.Int("port", 50052, "The server port")
	flag.Parse()

	db := api.InitClick()
	s := grpc.NewServer()
	pb.RegisterStatsServer(s, &api.Server{
		Click: db,
	})

	go api.ConsumeEvents(db)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
