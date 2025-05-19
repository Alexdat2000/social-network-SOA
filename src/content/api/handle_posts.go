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

func (s *Server) Get(_ context.Context, req *pb.UserPostRequest) (*pb.PostInfo, error) {
	var entry Entry
	err := s.EntriesDB.Where("id = ?", req.GetPostId()).First(&entry).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, status.Errorf(codes.NotFound, "post not found")
	} else if err != nil {
		log.Printf("Error when reading entry: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	if entry.Author != req.GetUser() && entry.IsPrivate == true {
		return nil, status.Error(codes.PermissionDenied, "no access to this private post")
	}

	ans := pb.PostInfo{
		PostId:       uint32(entry.ID),
		Title:        entry.Title,
		Description:  entry.Description,
		Author:       entry.Author,
		CreatedAt:    timestamppb.New(entry.CreatedAt),
		LastEditedAt: timestamppb.New(entry.LastEditedAt),
		IsPrivate:    entry.IsPrivate,
		Tags:         entry.Tags,
	}
	err = ReportGenericEventToKafka(s.Kafka, "post-views", req.GetUser(), req.GetPostId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &ans, nil
}

func (s *Server) Post(ctx context.Context, req *pb.PostRequest) (*pb.PostInfo, error) {
	t := time.Now()
	entry := Entry{
		Title:        req.GetTitle(),
		Description:  req.GetDescription(),
		Author:       req.GetUser(),
		CreatedAt:    t,
		LastEditedAt: t,
		IsPrivate:    req.GetIsPrivate(),
		Tags:         req.GetTags(),
	}
	if err := s.EntriesDB.Create(&entry).Error; err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return s.Get(ctx, &pb.UserPostRequest{
		User:   req.GetUser(),
		PostId: uint32(entry.ID),
	})
}

func (s *Server) Put(ctx context.Context, req *pb.PutRequest) (*pb.PostInfo, error) {
	updates := map[string]interface{}{
		"title":          req.GetTitle(),
		"description":    req.GetDescription(),
		"last_edited_at": time.Now(),
		"is_private":     req.GetIsPrivate(),
		"tags":           req.GetTags(),
	}

	result := s.EntriesDB.Model(&Entry{}).
		Where("id = ? AND author = ?", req.GetPostId(), req.GetUser()).
		Updates(updates)
	if result.Error != nil {
		return nil, status.Error(codes.Internal, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "post not found")
	}
	return s.Get(ctx, &pb.UserPostRequest{
		User:   req.GetUser(),
		PostId: req.GetPostId(),
	})
}

func (s *Server) Delete(_ context.Context, req *pb.UserPostRequest) (*emptypb.Empty, error) {
	result := s.EntriesDB.Where("id = ? AND author = ?", req.GetPostId(), req.GetUser()).Delete(&Entry{})
	if result.Error != nil {
		log.Printf("Error when deleting entry: %v", result.Error)
		return nil, status.Error(codes.Internal, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return &emptypb.Empty{}, status.Error(codes.NotFound, "post not found")
	} else {
		return &emptypb.Empty{}, nil
	}
}

func (s *Server) GetPosts(_ context.Context, req *pb.GetPostsRequest) (*pb.PostsInfo, error) {
	const pageSize = 2
	page := int(req.GetPage())
	if page <= 0 {
		return nil, status.Error(codes.InvalidArgument, "page must be greater than zero")
	}

	var totalCount int64
	if err := s.EntriesDB.Model(&Entry{}).Count(&totalCount).Error; err != nil {
		log.Printf("Error counting entries: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	var entries []Entry
	err := s.EntriesDB.Select("id").
		Order("id").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&entries).Error
	if err != nil {
		log.Printf("Error querying entries: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	postIds := make([]uint32, 0, len(entries))
	for _, entry := range entries {
		postIds = append(postIds, uint32(entry.ID))
	}
	return &pb.PostsInfo{
		TotalPages: uint32((totalCount + pageSize - 1) / pageSize),
		PostIds:    postIds,
	}, nil
}
