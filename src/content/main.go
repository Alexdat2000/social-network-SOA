package content

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "soa/content/content_grpc"
)

type server struct {
	pb.UnimplementedContentServer
}

func (s *server) Get(_ context.Context, req *pb.GetRequest) (*pb.PostInfo, error) {
	return &pb.PostInfo{}, nil
}

func (s *server) Post(_ context.Context, req *pb.PostRequest) (*pb.PostInfo, error) {
	return &pb.PostInfo{}, nil
}

func (s *server) Put(_ context.Context, req *pb.PutRequest) (*pb.PostInfo, error) {
	return &pb.PostInfo{}, nil
}

func main() {
	port := flag.Int("port", 50051, "The server port")
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterContentServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
