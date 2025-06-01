package api

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	pb "soa/stats/stats_grpc"
	"strings"
	"time"
)

func formatQuery(query string, metric pb.Metric) string {
	switch metric {
	case pb.Metric_VIEWS:
		return strings.Replace(query, "<table>", "stats.views", 1)
	case pb.Metric_LIKES:
		return strings.Replace(query, "<table>", "stats.likes", 1)
	case pb.Metric_COMMENTS:
		return strings.Replace(query, "<table>", "stats.comments", 1)
	}
	return ""
}

func (s *Server) Stats(ctx context.Context, req *pb.PostStatsRequest) (*pb.PostStats, error) {
	var views, likes, comments uint32
	query := `
		SELECT
			(SELECT count() FROM stats.views WHERE post_id = ?) AS views_count,
			(SELECT count() FROM stats.likes WHERE post_id = ?) AS likes_count,
			(SELECT count() FROM stats.comments WHERE post_id = ?) AS comments_count
	`
	err := s.Click.
		QueryRowContext(ctx, query, req.GetPostId(), req.GetPostId(), req.GetPostId()).
		Scan(&views, &likes, &comments)
	if err != nil {
		log.Printf("Query failed: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.PostStats{
		Views:    views,
		Likes:    likes,
		Comments: comments,
	}, nil
}

func (s *Server) Daily(ctx context.Context, req *pb.DailyRequest) (*pb.DailyStats, error) {
	query := formatQuery(`
        SELECT
            date,
            COUNT(*) AS rows_count
        FROM <table>
        GROUP BY date
        ORDER BY date ASC
    `, req.Metric)

	rows, err := s.Click.QueryContext(ctx, query)
	if err != nil {
		log.Printf("Query failed: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	defer rows.Close()

	var ans pb.DailyStats

	for rows.Next() {
		var date time.Time
		var count uint32
		if err := rows.Scan(&date, &count); err != nil {
			log.Printf("Failed to scan row: %v", err)
			return nil, status.Error(codes.Internal, err.Error())
		}
		ans.Stats = append(ans.Stats, &pb.DayStats{
			Date:  date.Format("2006-01-02"),
			Count: count,
		})
	}
	return &ans, nil
}

func (s *Server) TopPosts(ctx context.Context, req *pb.TopRequest) (*pb.TopPostList, error) {
	query := formatQuery(`
		SELECT
			post_id,
			COUNT(*) AS cnt
		FROM <table>
		GROUP BY post_id
		ORDER BY cnt DESC
		LIMIT 10
	`, req.Metric)

	rows, err := s.Click.QueryContext(ctx, query)
	if err != nil {
		log.Printf("Query failed: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	defer rows.Close()

	var ans pb.TopPostList
	for rows.Next() {
		var postId uint32
		var count uint32
		if err := rows.Scan(&postId, &count); err != nil {
			log.Printf("Failed to scan row: %v", err)
			return nil, status.Error(codes.Internal, err.Error())
		}
		ans.TopPosts = append(ans.TopPosts, &pb.PostInfo{
			PostId: postId,
			Count:  count,
		})
	}
	return &ans, nil
}

func (s *Server) TopUsers(ctx context.Context, req *pb.TopRequest) (*pb.TopUserList, error) {
	query := formatQuery(`
		SELECT
			author,
			COUNT(*) AS cnt
		FROM <table>
		GROUP BY author
		ORDER BY cnt DESC
		LIMIT 10
	`, req.Metric)

	rows, err := s.Click.QueryContext(ctx, query)
	if err != nil {
		log.Printf("Query failed: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	defer rows.Close()

	var ans pb.TopUserList
	for rows.Next() {
		var author string
		var count uint32
		if err := rows.Scan(&author, &count); err != nil {
			log.Printf("Failed to scan row: %v", err)
			return nil, status.Error(codes.Internal, err.Error())
		}
		ans.TopUsers = append(ans.TopUsers, &pb.UserInfo{
			Username: author,
			Count:    count,
		})
	}
	return &ans, nil
}
