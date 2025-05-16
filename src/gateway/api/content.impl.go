package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	pb "soa/gateway/content_grpc"
	"soa/gateway/utils"
	"strings"
	"time"
)

// Content

func (s Server) PostPosts(w http.ResponseWriter, r *http.Request) {
	ok, username := utils.Auth(w, r)
	if !ok {
		return
	}

	var req PostPostsJSONRequestBody
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	grpcReq := pb.PostRequest{
		User:        username,
		Title:       req.Title,
		Description: req.Content,
		IsPrivate:   false,
		Tags:        []string{},
	}
	if req.IsPrivate != nil {
		grpcReq.IsPrivate = *req.IsPrivate
	}
	if req.Tags != nil {
		grpcReq.Tags = *req.Tags
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ans, err := s.ContentAPI.Post(ctx, &grpcReq)
	if err != nil {
		log.Printf("Error POST grpc: %v", err)
		http.Error(w, "Error POST grpc: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	js, err := json.Marshal(ans)
	_, err = fmt.Fprintf(w, string(js))
	if err != nil {
		log.Printf("Error POST json: %v", err)
	}
}

func (s Server) GetPosts(w http.ResponseWriter, r *http.Request, params GetPostsParams) {
	//TODO implement me
	panic("implement me")
}

func (s Server) GetPostsPostId(w http.ResponseWriter, r *http.Request, postId int) {
	ok, username := utils.Auth(w, r)
	if !ok {
		return
	}

	if postId < 0 {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ans, err := s.ContentAPI.Get(ctx, &pb.UserPostRequest{
		User:   username,
		PostId: uint32(postId),
	})
	if errors.As(err, &sql.ErrNoRows) {
		http.Error(w, "Post not found", http.StatusNotFound)
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil && strings.Contains(err.Error(), "no access") {
		http.Error(w, "No access to post", http.StatusUnauthorized)
		return
	} else if err != nil {
		log.Printf("Error getting post: %v", err)
		http.Error(w, "Error getting post: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	js, _ := json.Marshal(ans)
	fmt.Fprintf(w, string(js))
}

func (s Server) PutPostsPostId(w http.ResponseWriter, r *http.Request, postId int) {
	//TODO implement me
	panic("implement me")
}

func (s Server) DeletePostsPostId(w http.ResponseWriter, r *http.Request, postId int) {
	//TODO implement me
	panic("implement me")
}
