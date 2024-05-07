package producer

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"go.uber.org/zap"
)

type IProducer interface {
	ProduceMessage(topic, value string) error
	Init() error
}

type Producer struct {
	zapLog      *zap.Logger
	kafkaConfig *kafka.ConfigMap
	producer    *kafka.Producer
}

func NewProducer(zapLog *zap.Logger, config *kafka.ConfigMap) *Producer {
	return &Producer{
		zapLog:      zapLog,
		kafkaConfig: config,
	}
}

func (p *Producer) Init() error {
	producer, err := kafka.NewProducer(p.kafkaConfig)
	if err != nil {
		p.zapLog.Error("Failed to create Kafka producer", zap.Error(err))
		return err
	}

	p.producer = producer
	log.Println("Kafka producer created")
	return nil
}

func (p *Producer) ProduceMessage(topic string, value []byte) error {
	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          value,
	}

	err := p.producer.Produce(msg, nil)
	if err == nil {
		p.zapLog.Info("Message sent to Kafka", zap.String("Topic", topic), zap.ByteString("Value", value))
	}
	return err
}
