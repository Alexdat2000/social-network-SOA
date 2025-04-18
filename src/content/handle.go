package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"soa/content/api"
	pb "soa/content/content_grpc"
	"strings"
	"time"
)

var (
	noAccessError = errors.New("no access")
)

func (s *server) Get(_ context.Context, req *pb.UserPostRequest) (*pb.PostInfo, error) {
	log.Println("Received GET request")
	re := fmt.Sprintf(`select * from entries where id='%d'`, req.GetPostId())
	row := api.DB.QueryRow(re)

	var ans pb.PostInfo
	var tags string
	var createdAt time.Time
	var lastEditedAt time.Time
	err := row.Scan(&ans.PostId, &ans.Title, &ans.Description, &ans.Author, &createdAt, &lastEditedAt, &ans.IsPrivate, &tags)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, err
	} else if err != nil {
		log.Printf("Error when reading row: %v", err)
		return nil, err
	}
	if ans.GetAuthor() != req.GetUser() && ans.IsPrivate {
		return nil, noAccessError
	}
	ans.Tags = strings.Split(tags, ",")
	ans.CreatedAt = timestamppb.New(createdAt)
	ans.LastEditedAt = timestamppb.New(lastEditedAt)
	return &ans, nil
}

func (s *server) Post(ctx context.Context, req *pb.PostRequest) (*pb.PostInfo, error) {
	log.Println("Received POST request")
	id, err := uuid.NewUUID()
	if err != nil {
		log.Printf("Error when generating uuid: %v", err)
		return nil, err
	}
	idInt := id.ID() / 2
	t := time.Now()
	_, err = api.DB.Exec(`insert into entries (id, title, description, author, created_at, last_edited_at, is_private, tags)
values ($1, $2, $3, $4, $5, $6, $7, $8)`,
		idInt, req.GetTitle(), req.GetDescription(), req.GetUser(), t, t, req.GetIsPrivate(), strings.Join(req.GetTags(), ","))
	if err != nil {
		return nil, err
	}
	return s.Get(ctx, &pb.UserPostRequest{
		User:   req.GetUser(),
		PostId: idInt,
	})
}

func (s *server) Put(ctx context.Context, req *pb.PutRequest) (*pb.PostInfo, error) {
	log.Println("Received PUT request")
	res, err := api.DB.Exec(`update entries
set title = $3, 
description = $4,
last_edited_at = $5,
is_private = $6,
tags = $7
where id = $1 and author = $2`,
		req.GetPostId(),
		req.GetUser(),
		req.GetTitle(),
		req.GetDescription(),
		time.Now(),
		req.GetIsPrivate(),
		strings.Join(req.GetTags(), ","))

	if err != nil {
		log.Printf("Error when updating entries: %v", err)
		return nil, err
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return nil, errors.New("not found")
	}
	return s.Get(ctx, &pb.UserPostRequest{
		User:   req.GetUser(),
		PostId: req.GetPostId(),
	})
}

func (s *server) Delete(_ context.Context, req *pb.UserPostRequest) (*pb.DeleteResult, error) {
	log.Println("Received DELETE request")
	res, err := api.DB.Exec(`delete from entries where id = $1 and author = $2`, req.GetPostId(), req.GetUser())
	if err != nil {
		log.Printf("Error when deleting entry: %v", err)
		return nil, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error when deleting entry: %v", err)
		return nil, err
	}
	if count == 0 {
		return &pb.DeleteResult{Successful: false}, nil
	} else {
		return &pb.DeleteResult{Successful: true}, nil
	}
}

func (s *server) GetPosts(_ context.Context, req *pb.GetPostsRequest) (*pb.PostsInfo, error) {
	log.Println("Received GetAll request")
	pageSize := 2

	rows, err := api.DB.Query(`select id from entries order by id offset $1 limit $2`, int(req.GetPage()-1)*pageSize, pageSize)
	if err != nil {
		log.Printf("Error when querying entries: %v", err)
		return &pb.PostsInfo{}, err
	}
	defer rows.Close()

	var ids []uint32
	for rows.Next() {
		var id uint32
		err := rows.Scan(&id)
		if err == nil {
			ids = append(ids, id)
		} else {
			log.Printf("Entry error: %v", err)
		}
	}

	var totalCount int
	err = api.DB.QueryRow("SELECT COUNT(*) FROM entries").Scan(&totalCount)
	if err != nil {
		log.Printf("Error when querying entries: %v", err)
		return nil, err
	}
	return &pb.PostsInfo{
		Page:       req.GetPage(),
		TotalPages: uint32((totalCount + pageSize - 1) / pageSize),
		PostIds:    ids,
	}, nil
}
