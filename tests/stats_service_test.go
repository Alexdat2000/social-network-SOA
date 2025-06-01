package main

import (
	"context"
	"github.com/stretchr/testify/assert"
	"soa/tests/stats_grpc"
	"testing"
)

func prepareClick() {
	ExecuteInClick([]string{`ALTER TABLE stats.views DELETE WHERE 1=1;`,
		`INSERT INTO stats.views (post_id, author, date) VALUES
(1, 'Alex', '2025-05-24'),
(2, 'Alex2', '2025-05-24'),
(1, 'Alex', '2025-05-25');
`})

	ExecuteInClick([]string{`ALTER TABLE stats.likes DELETE WHERE 1=1;`,
		`INSERT INTO stats.likes (post_id, author, date) VALUES
(1, 'Alex', '2025-05-24'),
(2, 'Alex2', '2025-05-25'),
(1, 'Alex', '2025-05-25');
`})

	ExecuteInClick([]string{`ALTER TABLE stats.comments DELETE WHERE 1=1;`,
		`INSERT INTO stats.comments (post_id, author, date) VALUES
(1, 'Alex', '2025-05-24'),
(1, 'Alex', '2025-05-26'),
(1, 'Alex', '2025-05-27');
`})
}

func TestPostStats(t *testing.T) {
	prepareClick()

	grpc := stats_grpc.InitStatsClient("stats:50052")
	stat, _ := grpc.Stats(context.Background(), &stats_grpc.PostStatsRequest{PostId: 1})
	assert.Equal(t, uint32(2), stat.Views)
	assert.Equal(t, uint32(2), stat.Likes)
	assert.Equal(t, uint32(3), stat.Comments)
}

func TestDailyStats(t *testing.T) {
	prepareClick()

	grpc := stats_grpc.InitStatsClient("stats:50052")
	stat, _ := grpc.Daily(context.Background(), &stats_grpc.DailyRequest{PostId: 1, Metric: stats_grpc.Metric_LIKES})
	assert.Equal(t, "2025-05-24", stat.Stats[0].Date)
	assert.Equal(t, uint32(1), stat.Stats[0].Count)
	assert.Equal(t, "2025-05-25", stat.Stats[1].Date)
	assert.Equal(t, uint32(2), stat.Stats[1].Count)
}

func TestTopUsers(t *testing.T) {
	prepareClick()

	grpc := stats_grpc.InitStatsClient("stats:50052")
	stat, _ := grpc.TopUsers(context.Background(), &stats_grpc.TopRequest{Metric: stats_grpc.Metric_VIEWS})
	assert.Equal(t, "Alex", stat.TopUsers[0].Username)
	assert.Equal(t, uint32(2), stat.TopUsers[0].Count)
	assert.Equal(t, "Alex2", stat.TopUsers[1].Username)
	assert.Equal(t, uint32(1), stat.TopUsers[1].Count)
}
