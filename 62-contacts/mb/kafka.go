package mb

import (
	"context"

	kafka "github.com/segmentio/kafka-go"
)

type Kafka struct {
	Conns []string
}

func (k *Kafka) WriteMessage(message []byte, topic string) error {
	// make a writer that produces to topic-A, using the least-bytes distribution
	w := &kafka.Writer{
		Addr:     kafka.TCP(k.Conns...), //kafka.TCP("localhost:9092", "localhost:9093", "localhost:9094"),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			//Key:   []byte("Key-A"),
			Value: message,
		},
	)
	if err != nil {
		return err
	}

	if err := w.Close(); err != nil {
		return err
	}
	return nil
}
