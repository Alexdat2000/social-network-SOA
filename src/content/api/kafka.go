package api

import (
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"time"
)

func InitKafka() *kafka.Producer {
	kafkaBroker := "kafka:9092"
	kaf, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": kafkaBroker})
	if err != nil {
		log.Fatalf("Failed to create producer: %v", err)
		return nil
	}
	log.Println("Successfully connected to the Kafka broker!")
	return kaf
}

func reportToKafka(kaf *kafka.Producer, topic string, value []byte) error {
	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          value,
	}

	deliveryChan := make(chan kafka.Event)
	err := kaf.Produce(message, deliveryChan)
	if err != nil {
		return err
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)
	close(deliveryChan)
	return m.TopicPartition.Error
}

type GetEvent struct {
	Username  string `json:"username"`
	PostId    uint32 `json:"post_id"`
	Timestamp string `json:"timestamp"`
}

func ReportGenericEventToKafka(kaf *kafka.Producer, topic, user string, postId uint32) error {
	msg, _ := json.Marshal(GetEvent{user, postId, time.Now().Format(time.RFC3339)})
	err := reportToKafka(kaf, topic, msg)
	if err != nil {
		log.Printf("Error reporting GET to kafka: %v", err)
		return err
	} else {
		return nil
	}
}
