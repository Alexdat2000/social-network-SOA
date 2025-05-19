package api

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"gorm.io/gorm"
	pb "soa/content/content_grpc"
	"time"
)

type Server struct {
	pb.UnimplementedContentServer
	EntriesDB  *gorm.DB
	CommentsDB *gorm.DB
	Kafka      *kafka.Producer
}

type Entry struct {
	ID           uint      `gorm:"primaryKey;autoIncrement;column:id"`
	Title        string    `gorm:"type:text;not null;column:title"`
	Description  string    `gorm:"type:text;not null;column:description"`
	Author       string    `gorm:"type:varchar(32);not null;column:author"`
	CreatedAt    time.Time `gorm:"type:timestamptz;not null;default:now();column:created_at"`
	LastEditedAt time.Time `gorm:"type:timestamptz;not null;default:now();column:last_edited_at"`
	IsPrivate    bool      `gorm:"not null;default:true;column:is_private"`
	Tags         []string  `gorm:"type:text[];not null;default:'{}';column:tags"`
}

type Comment struct {
	ID        int       `gorm:"primaryKey;autoIncrement;column:id"`
	PostID    int       `gorm:"column:post_id"`
	Author    string    `gorm:"type:varchar(32);not null;column:author"`
	Text      string    `gorm:"type:text;not null;column:text"`
	CreatedAt time.Time `gorm:"type:timestamp;not null;column:created_at"`
}
