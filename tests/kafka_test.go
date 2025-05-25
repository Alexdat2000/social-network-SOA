package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestKafkaViewsTopic(t *testing.T) {
	ClearTableClick("stats.views")

	kafka := InitKafka()
	assert.Equal(t, 0, CalcRowsInClick("stats.views"))
	_ = ReportGenericEventToKafka(kafka, "post-views", 4, "Alex")
	_ = ReportGenericEventToKafka(kafka, "post-views", 5, "Alex2")
	time.Sleep(100 * time.Millisecond)
	assert.Equal(t, 2, CalcRowsInClick("stats.views"))
}

func TestKafkaCommentsTopic(t *testing.T) {
	ClearTableClick("stats.comments")

	kafka := InitKafka()
	assert.Equal(t, 0, CalcRowsInClick("stats.comments"))
	_ = ReportGenericEventToKafka(kafka, "post-comments", 1, "Alex")
	_ = ReportGenericEventToKafka(kafka, "post-comments", 1, "Alex")
	time.Sleep(100 * time.Millisecond)
	assert.Equal(t, 2, CalcRowsInClick("stats.comments"))
}
