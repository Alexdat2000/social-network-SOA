package api

import (
	"context"
	"encoding/json"
	"log"
	"math"
	"net/http"
	pb "soa/gateway/content_grpc"
	"soa/gateway/utils"
	"time"
)

// Interactions

func (s Server) PostPostsPostIdLikes(w http.ResponseWriter, r *http.Request, postId int) {
	if !(postId >= 0 && postId <= math.MaxUint32) {
		http.Error(w, "PostId is invalid", http.StatusNotFound)
	}

	ok, username := utils.Auth(w, r)
	if !ok {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ans, err := s.ContentAPI.LikePost(ctx, &pb.UserPostRequest{
		User:   username,
		PostId: uint32(postId),
	})
	if err != nil {
		translateGrpcErrorToHttp(err, w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(ans); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}

func (s Server) PostPostsPostIdComments(w http.ResponseWriter, r *http.Request, postId int) {
	if !(postId >= 0 && postId <= math.MaxUint32) {
		http.Error(w, "PostId is invalid", http.StatusNotFound)
	}

	ok, username := utils.Auth(w, r)
	if !ok {
		return
	}

	var req PostPostsPostIdCommentsJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ans, err := s.ContentAPI.PostComment(ctx, &pb.PostCommentRequest{
		User:   username,
		PostId: uint32(postId),
		Text:   req.Text,
	})
	if err != nil {
		translateGrpcErrorToHttp(err, w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(ans); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}

func (s Server) GetPostsPostIdComments(w http.ResponseWriter, r *http.Request, postId int, params GetPostsPostIdCommentsParams) {
	if params.Page < 0 {
		http.Error(w, "Page is invalid", http.StatusBadRequest)
		return
	}

	ok, _ := utils.Auth(w, r)
	if !ok {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	ans, err := s.ContentAPI.GetComments(ctx, &pb.GetCommentsRequest{
		PostId: uint32(postId),
		Page:   uint32(params.Page),
	})
	if err != nil {
		translateGrpcErrorToHttp(err, w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(ans); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}
