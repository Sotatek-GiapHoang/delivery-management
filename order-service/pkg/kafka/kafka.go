package kafka

import (
	"context"
	"encoding/json"
	"order-service/internal/models"
	"strconv"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(brokers []string, topic string) *Producer {
	w := &kafka.Writer{
		Addr:     kafka.TCP(brokers...),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
	return &Producer{writer: w}
}

func (p *Producer) SendOrderCreatedEvent(order *models.OrderCreatedEvent) error {
	orderJSON, err := json.Marshal(order)
	if err != nil {
		return err
	}

	message := kafka.Message{
		Key:   []byte(strconv.FormatUint(uint64(order.ID), 10)),
		Value: orderJSON,
	}

	return p.writer.WriteMessages(context.Background(), message)
}

func (p *Producer) Close() error {
	return p.writer.Close()
}
