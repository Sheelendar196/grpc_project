package main

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	logger "golang.org/x/exp/slog"
)

var consumer *kafka.Consumer

func main() {
	topics := []string{"employee_data"}
	createKafkaConsumer()
	ConsumeData(topics)
}
func createKafkaConsumer() (*kafka.Consumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  "localhost:9092",
		"group.id":           "console-consumer-21564",
		"session.timeout.ms": 6000,
		"auto.offset.reset":  "earliest",
	})
	if err != nil {
		panic(err)
	}
	consumer = c
	return c, err
}

func ConsumeData(topics []string) {
	consumer.SubscribeTopics(topics, nil)
	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			logger.Error("get error consume data err: %v", err)
		} else {
			logger.Info("message received from topic data: %v", string(msg.Value))
		}
	}
}
