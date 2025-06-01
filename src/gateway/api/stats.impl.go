package api

import (
	"context"
	"encoding/json"
	"log"
	"math"
	"net/http"
	pb "soa/gateway/stats_grpc"
	"soa/gateway/utils"
	"time"
)

// Statistics

func (s Server) GetPostsPostIdStats(w http.ResponseWriter, r *http.Request, postId int) {
	if !(postId >= 0 && postId <= math.MaxUint32) {
		http.Error(w, "PostId is invalid", http.StatusNotFound)
	}

	ok, _ := utils.Auth(w, r)
	if !ok {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ans, err := s.StatsAPI.Stats(ctx, &pb.PostStatsRequest{
		PostId: uint32(postId),
	})
	if err != nil {
		utils.TranslateGrpcErrorToHttp(err, w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(ans); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}

func (s Server) GetPostsPostIdStatsDaily(w http.ResponseWriter, r *http.Request, postId int, params GetPostsPostIdStatsDailyParams) {
	if !(postId >= 0 && postId <= math.MaxUint32) {
		http.Error(w, "PostId is invalid", http.StatusNotFound)
	}

	ok, _ := utils.Auth(w, r)
	if !ok {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var metric pb.Metric
	switch params.Metric {
	case GetPostsPostIdStatsDailyParamsMetricViews:
		metric = pb.Metric_VIEWS
	case GetPostsPostIdStatsDailyParamsMetricLikes:
		metric = pb.Metric_LIKES
	case GetPostsPostIdStatsDailyParamsMetricComments:
		metric = pb.Metric_COMMENTS
	default:
		http.Error(w, "Metric not supported", http.StatusNotFound)
	}

	ans, err := s.StatsAPI.Daily(ctx, &pb.DailyRequest{
		PostId: uint32(postId),
		Metric: metric,
	})
	if err != nil {
		utils.TranslateGrpcErrorToHttp(err, w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(ans.GetStats()); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}

func (s Server) GetPostsTop10(w http.ResponseWriter, r *http.Request, params GetPostsTop10Params) {
	ok, _ := utils.Auth(w, r)
	if !ok {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var metric pb.Metric
	switch params.Metric {
	case GetPostsTop10ParamsMetricViews:
		metric = pb.Metric_VIEWS
	case GetPostsTop10ParamsMetricLikes:
		metric = pb.Metric_LIKES
	case GetPostsTop10ParamsMetricComments:
		metric = pb.Metric_COMMENTS
	default:
		http.Error(w, "Metric not supported", http.StatusNotFound)
	}

	ans, err := s.StatsAPI.TopPosts(ctx, &pb.TopRequest{
		Metric: metric,
	})
	if err != nil {
		utils.TranslateGrpcErrorToHttp(err, w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(ans.GetTopPosts()); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}

func (s Server) GetUsersTop10(w http.ResponseWriter, r *http.Request, params GetUsersTop10Params) {
	ok, _ := utils.Auth(w, r)
	if !ok {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var metric pb.Metric
	switch params.Metric {
	case Views:
		metric = pb.Metric_VIEWS
	case Likes:
		metric = pb.Metric_LIKES
	case Comments:
		metric = pb.Metric_COMMENTS
	default:
		http.Error(w, "Metric not supported", http.StatusNotFound)
	}

	ans, err := s.StatsAPI.TopUsers(ctx, &pb.TopRequest{
		Metric: metric,
	})
	if err != nil {
		utils.TranslateGrpcErrorToHttp(err, w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(ans.GetTopUsers()); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}
