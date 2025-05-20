package api

import (
	"context"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func initTopics() {
	admin, err := kafka.NewAdminClient(&kafka.ConfigMap{"bootstrap.servers": "kafka:9092"})
	if err != nil {
		panic(err)
	}
	defer admin.Close()

	topics := []kafka.TopicSpecification{
		{Topic: "post-views", NumPartitions: 3, ReplicationFactor: 1},
		{Topic: "post-likes", NumPartitions: 3, ReplicationFactor: 1},
		{Topic: "post-comments", NumPartitions: 3, ReplicationFactor: 1},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	results, err := admin.CreateTopics(ctx, topics)
	if err != nil {
		panic(err)
	}

	for _, result := range results {
		if result.Error.Code() == kafka.ErrTopicAlreadyExists {
			fmt.Printf("Topic %s already exists\n", result.Topic)
		} else if result.Error.Code() != kafka.ErrNoError {
			fmt.Printf("Failed to create topic %s: %v\n", result.Topic, result.Error)
		} else {
			fmt.Printf("Topic %s created successfully\n", result.Topic)
		}
	}
}

func ConsumeEvents() {
	initTopics()

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"group.id":          "consumer",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	err = consumer.SubscribeTopics([]string{"post-views", "post-likes", "post-comments"}, nil)
	if err != nil {
		panic(err)
	}

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Kafka consumer started")

	run := true
	for run {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev := consumer.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				fmt.Printf("Message on %s: %s\n", e.TopicPartition, string(e.Value))
				if e.Headers != nil {
					fmt.Printf("Headers: %v\n", e.Headers)
				}
			case kafka.Error:
				fmt.Fprintf(os.Stderr, "Error: %v\n", e)
				if e.Code() == kafka.ErrAllBrokersDown {
					run = false
				}
			default:
				// Ignore other events
			}
		}
	}
}
