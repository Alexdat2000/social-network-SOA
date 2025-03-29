package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log"
	"soa/content/api"
	pb "soa/content/content_grpc"
	"strings"
	"time"
)

func (s *server) Get(_ context.Context, req *pb.GetRequest) (*pb.PostInfo, error) {
	re := fmt.Sprintf(`select * from entries where id='%d'`, req.GetPostId())
	row := api.DB.QueryRow(re)

	var ans pb.PostInfo
	var tags string
	err := row.Scan(&ans.PostId, &ans.Title, &ans.Description, &ans.Author, &ans.CreatedAt, &ans.LastEditedAt, &ans.IsPrivate, &tags)
	if errors.Is(err, sql.ErrNoRows) {
		return &pb.PostInfo{}, err
	} else if err != nil {
		log.Printf("Error when reading row: %v", err)
		return &pb.PostInfo{}, err
	}
	if ans.GetAuthor() != req.GetUser() && ans.IsPrivate {
		return &pb.PostInfo{}, errors.New("no access")
	}
	ans.Tags = strings.Split(tags, ",")
	return &ans, nil
}

func (s *server) Post(ctx context.Context, req *pb.PostRequest) (*pb.PostInfo, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		log.Printf("Error when generating uuid: %v", err)
		return &pb.PostInfo{}, err
	}
	idInt := id.ID() / 2
	t := time.Now()
	_, err = api.DB.Exec(`insert into entries (id, title, description, author, created_at, last_edited_at, is_private, tags)
values ($1, $2, $3, $4, $5, $6, $7, $8)`,
		idInt, req.GetTitle(), req.GetDescription(), req.GetUser(), t, t, req.GetIsPrivate(), strings.Join(req.GetTags(), ","))
	if err != nil {
		return &pb.PostInfo{}, err
	}
	return s.Get(ctx, &pb.GetRequest{
		User:   req.GetUser(),
		PostId: idInt,
	})
}

func (s *server) Put(ctx context.Context, req *pb.PutRequest) (*pb.PostInfo, error) {
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
		return &pb.PostInfo{}, err
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return &pb.PostInfo{}, errors.New("not found")
	}
	return s.Get(ctx, &pb.GetRequest{
		User:   req.GetUser(),
		PostId: req.GetPostId(),
	})
}
