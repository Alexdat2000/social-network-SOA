package main

import (
	"context"
	"github.com/stretchr/testify/assert"
	"soa/tests/content_grpc"
	"testing"
)

func TestPostContentService(t *testing.T) {
	ClearTablePostgres("postgres_content", 5433, "content", "entries")

	grpc := content_grpc.InitContentClient("content:50051")
	_, _ = grpc.Post(context.Background(), &content_grpc.PostRequest{
		User:        "Alex",
		Title:       "Test post",
		Description: "Test post",
	})
	assert.Equal(t, 1, CalcRowsInTable("postgres_content", 5433, "content", "entries"))

	post, _ := grpc.Get(context.Background(), &content_grpc.UserPostRequest{
		User:   "Alex",
		PostId: 1,
	})
	assert.Equal(t, "Test post", post.Title)
}

func TestPostModificationsContentService(t *testing.T) {
	ClearTablePostgres("postgres_content", 5433, "content", "entries")

	grpc := content_grpc.InitContentClient("content:50051")
	_, _ = grpc.Post(context.Background(), &content_grpc.PostRequest{
		User:        "Alex",
		Title:       "Test post",
		Description: "Test post",
	})
	assert.Equal(t, 1, CalcRowsInTable("postgres_content", 5433, "content", "entries"))

	_, _ = grpc.Put(context.Background(), &content_grpc.PutRequest{
		User:   "Alex",
		PostId: 1,
		Title:  "New title",
	})
	post, _ := grpc.Get(context.Background(), &content_grpc.UserPostRequest{
		User:   "Alex",
		PostId: 1,
	})
	assert.Equal(t, "New title", post.Title)

	_, _ = grpc.Delete(context.Background(), &content_grpc.UserPostRequest{
		User:   "Alex",
		PostId: 1,
	})
	assert.Equal(t, 0, CalcRowsInTable("postgres_content", 5433, "content", "entries"))
}

func TestPostListContentService(t *testing.T) {
	ClearTablePostgres("postgres_content", 5433, "content", "entries")

	grpc := content_grpc.InitContentClient("content:50051")
	_, _ = grpc.Post(context.Background(), &content_grpc.PostRequest{
		User:        "Alex",
		Title:       "Test post 1",
		Description: "Test post 1",
	})
	_, _ = grpc.Post(context.Background(), &content_grpc.PostRequest{
		User:        "Alex2",
		Title:       "Test post 2",
		Description: "Test post 2",
	})
	_, _ = grpc.Post(context.Background(), &content_grpc.PostRequest{
		User:        "Alex",
		Title:       "Test post 3",
		Description: "Test post 3",
	})
	assert.Equal(t, 3, CalcRowsInTable("postgres_content", 5433, "content", "entries"))

	post, _ := grpc.GetPosts(context.Background(), &content_grpc.GetPostsRequest{Page: 1})
	assert.Equal(t, uint32(2), post.TotalPages)
	assert.Equal(t, uint32(1), post.PostIds[0])
	assert.Equal(t, uint32(2), post.PostIds[1])
}
