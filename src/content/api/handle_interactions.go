package api

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	pb "soa/content/content_grpc"
	"time"
)

func (s *Server) LikePost(_ context.Context, req *pb.UserPostRequest) (*emptypb.Empty, error) {
	err := ReportGenericEventToKafka(s.Kafka, "post-likes", req.GetUser(), req.GetPostId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	} else {
		return &emptypb.Empty{}, nil
	}
}

func (s *Server) PostComment(_ context.Context, req *pb.PostCommentRequest) (*emptypb.Empty, error) {
	entry := Comment{
		PostID:    int(req.PostId),
		Author:    req.User,
		Text:      req.Text,
		CreatedAt: time.Now(),
	}
	if err := s.DB.Create(&entry).Error; err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	err := ReportGenericEventToKafka(s.Kafka, "post-comments", req.GetUser(), req.GetPostId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	} else {
		return &emptypb.Empty{}, nil
	}
}

func (s *Server) GetComments(_ context.Context, req *pb.GetCommentsRequest) (*pb.CommentsInfo, error) {
	const pageSize = 2
	page := int(req.GetPage())
	if page <= 0 {
		return nil, status.Error(codes.InvalidArgument, "page must be greater than zero")
	}

	var totalCount int64
	if err := s.DB.Model(&Comment{}).Where("post_id = ?", req.GetPostId()).Count(&totalCount).Error; err != nil {
		log.Printf("Error counting comments: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	var comments []Comment
	err := s.DB.
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&comments).Error
	if err != nil {
		log.Printf("Error querying entries: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	var commentInfos []*pb.CommentInfo
	for _, comment := range comments {
		commentInfos = append(commentInfos, &pb.CommentInfo{
			Id:        uint32(comment.ID),
			Author:    comment.Author,
			Text:      comment.Text,
			CreatedAt: timestamppb.New(comment.CreatedAt),
		})
	}

	return &pb.CommentsInfo{
		TotalPages: uint32(totalCount),
		Comments:   commentInfos,
	}, nil
}
