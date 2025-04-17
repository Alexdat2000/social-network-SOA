package api

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"time"
)

var KAFKA *kafka.Producer

func ConnectToKafka() {
	kafkaBroker := "kafka:9092"
	var err error
	KAFKA, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": kafkaBroker})
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		return
	}
	log.Println("Successfully connected to the Kafka broker!")
}

type RegisterEvent struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Timestamp string `json:"timestamp"`
}

func ReportRegisterToKafka(username, email string, t time.Time) error {
	topic := "user-registrations"

	msg, _ := json.Marshal(RegisterEvent{username, email, t.Format(time.RFC3339)})
	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          msg,
	}

	deliveryChan := make(chan kafka.Event)
	err := KAFKA.Produce(message, deliveryChan)
	if err != nil {
		return err
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)
	close(deliveryChan)
	return m.TopicPartition.Error
}
