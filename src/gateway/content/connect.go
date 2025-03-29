package content

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"

	pb "soa/gateway/content_grpc"
)

var c pb.ContentClient

func InitGrpc(addr string) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c = pb.NewContentClient(conn)
}
