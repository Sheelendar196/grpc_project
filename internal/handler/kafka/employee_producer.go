package kafka

import (
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sheelendar196/go-projects/grpc_project/proto"

	logger "golang.org/x/exp/slog"
)

type EmployeeProducer struct {
	producer *kafka.Producer
	topic    string
}

func (ep *EmployeeProducer) PushEmployeeTokafka(emp *proto.Employee) error {
	empBytes, err := json.Marshal(emp)
	if err != nil {
		logger.Error("marshling error : %v", err)
	}
	return ep.ProduceData(empBytes)
}

// topic = "employee_data"
func (ep *EmployeeProducer) ProduceData(msgBytes []byte) error {

	m := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &ep.topic, Partition: kafka.PartitionAny},
		Value:          msgBytes,
	}
	err := ep.producer.Produce(m, nil)
	if err != nil {
		logger.Error("failed to produce message, topic: %v, error: %v, key: %v", ep.topic, err)
	} else {
		logger.Info("message has beed pushed to kafka successfully")
	}
	return err
}

func CreateKafkaProducer(topic string) (*EmployeeProducer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	})
	if err != nil {
		panic(err)
	}
	empKafka := &EmployeeProducer{producer: p, topic: topic}
	return empKafka, err
}
