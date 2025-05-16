package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"soa/content/api"
	pb "soa/content/content_grpc"
)

func main() {
	port := flag.Int("port", 50051, "The server port")
	flag.Parse()

	s := grpc.NewServer()
	pb.RegisterContentServer(s, &api.Server{
		Db:    api.InitDB(),
		Kafka: api.InitKafka(),
	})

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
