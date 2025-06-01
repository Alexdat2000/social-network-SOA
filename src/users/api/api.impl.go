package api

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"gorm.io/gorm"
)

type Server struct {
	DB       *gorm.DB
	Kafka    *kafka.Producer
	Handlers *AuthHandlers
}
