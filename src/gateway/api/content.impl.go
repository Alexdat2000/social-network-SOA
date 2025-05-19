package api

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"math"
	"net/http"
	pb "soa/gateway/content_grpc"
	"soa/gateway/utils"
	"time"
)

func translateGrpcErrorToHttp(err error, w http.ResponseWriter) {
	st, ok := status.FromError(err)
	if !ok {
		st = status.New(codes.Internal, err.Error())
	}
	if st.Code() == codes.Internal {
		log.Printf("Error grpc: %v", err)
	}
	http.Error(w, st.Message(), utils.GrpcCodeToHTTPStatus(st.Code()))
}

// Content

func (s Server) PostPosts(w http.ResponseWriter, r *http.Request) {
	ok, username := utils.Auth(w, r)
	if !ok {
		return
	}

	var req PostPostsJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
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
		translateGrpcErrorToHttp(err, w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(ans); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}

func (s Server) GetPostsPostId(w http.ResponseWriter, r *http.Request, postId int) {
	if !(postId >= 0 && postId <= math.MaxUint32) {
		http.Error(w, "PostId is invalid", http.StatusNotFound)
	}

	ok, username := utils.Auth(w, r)
	if !ok {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ans, err := s.ContentAPI.Get(ctx, &pb.UserPostRequest{
		User:   username,
		PostId: uint32(postId),
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

func (s Server) PutPostsPostId(w http.ResponseWriter, r *http.Request, postId int) {
	if !(postId >= 0 && postId <= math.MaxUint32) {
		http.Error(w, "PostId is invalid", http.StatusNotFound)
		return
	}
	ok, username := utils.Auth(w, r)
	if !ok {
		return
	}

	var req PutPostsPostIdJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	grpcReq := pb.PutRequest{
		User:   username,
		PostId: uint32(postId),
	}
	if req.Title != nil {
		grpcReq.Title = *req.Title
	}
	if req.Content != nil {
		grpcReq.Description = *req.Content
	}
	if req.IsPrivate != nil {
		grpcReq.IsPrivate = *req.IsPrivate
	}
	if req.Tags != nil {
		grpcReq.Tags = *req.Tags
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	ans, err := s.ContentAPI.Put(ctx, &grpcReq)
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

func (s Server) DeletePostsPostId(w http.ResponseWriter, r *http.Request, postId int) {
	if !(postId >= 0 && postId <= math.MaxUint32) {
		http.Error(w, "PostId is invalid", http.StatusNotFound)
	}

	ok, username := utils.Auth(w, r)
	if !ok {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ans, err := s.ContentAPI.Delete(ctx, &pb.UserPostRequest{
		User:   username,
		PostId: uint32(postId),
	})
	if err != nil {
		translateGrpcErrorToHttp(err, w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	if err := json.NewEncoder(w).Encode(ans); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}

func (s Server) GetPosts(w http.ResponseWriter, r *http.Request, params GetPostsParams) {
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

	ans, err := s.ContentAPI.GetPosts(ctx, &pb.GetPostsRequest{
		Page: uint32(params.Page),
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
