package api

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"log"
	pb "soa/content/content_grpc"
	"time"
)

func checkPostAccess(db *gorm.DB, postId uint32, user string) (string, error) {
	var entry Entry
	err := db.Where("id = ?", postId).First(&entry).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", status.Errorf(codes.NotFound, "post not found")
	} else if err != nil {
		log.Printf("Error when reading entry: %v", err)
		return "", status.Error(codes.Internal, err.Error())
	}
	if entry.Author != user && *entry.IsPrivate == true {
		return "", status.Error(codes.PermissionDenied, "no access to this private post")
	}
	return entry.Author, nil
}

func (s *Server) LikePost(_ context.Context, req *pb.UserPostRequest) (*emptypb.Empty, error) {
	author, err := checkPostAccess(s.DB, req.GetPostId(), req.GetUser())
	if err != nil {
		return nil, err
	}

	err = ReportGenericEventToKafka(s.Kafka, "post-likes", req.GetPostId(), author)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	} else {
		return &emptypb.Empty{}, nil
	}
}

func (s *Server) PostComment(_ context.Context, req *pb.PostCommentRequest) (*emptypb.Empty, error) {
	author, err := checkPostAccess(s.DB, req.GetPostId(), req.GetUser())
	if err != nil {
		return nil, err
	}

	entry := Comment{
		PostID:    int(req.PostId),
		Author:    req.User,
		Text:      req.Text,
		CreatedAt: time.Now(),
	}
	if err := s.DB.Create(&entry).Error; err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = ReportGenericEventToKafka(s.Kafka, "post-comments", req.GetPostId(), author)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	} else {
		return &emptypb.Empty{}, nil
	}
}

func (s *Server) GetComments(_ context.Context, req *pb.GetCommentsRequest) (*pb.CommentsInfo, error) {
	_, err := checkPostAccess(s.DB, req.GetPostId(), req.GetUser())
	if err != nil {
		return nil, err
	}

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
	err = s.DB.
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
		TotalPages: uint32((totalCount + pageSize - 1) / pageSize),
		Comments:   commentInfos,
	}, nil
}
