package kafka

import (
	"context"
	"delivery-service/internal/models"
	"encoding/json"
	"log"
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

func (p *Producer) SendDeliveryCreatedEvent(delivery *models.Delivery) error {
	deliveryJSON, err := json.Marshal(delivery)
	if err != nil {
		return err
	}

	message := kafka.Message{
		Key:   []byte(strconv.FormatUint(uint64(delivery.ID), 10)),
		Value: deliveryJSON,
	}

	return p.writer.WriteMessages(context.Background(), message)
}

func (p *Producer) Close() error {
	return p.writer.Close()
}

type Consumer struct {
	reader *kafka.Reader
}

func NewConsumer(brokers []string, topic, groupID string) *Consumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		GroupID:  groupID,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
	return &Consumer{reader: r}
}

func (c *Consumer) ConsumeOrderCreatedEvents(handler func([]byte) error) {
	for {
		m, err := c.reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error reading message: %v", err)
			continue
		}
		if err := handler(m.Value); err != nil {
			log.Printf("Error handling message: %v", err)
		}
	}
}

func (c *Consumer) Close() error {
	return c.reader.Close()
}
