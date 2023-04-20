package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	kafka "github.com/segmentio/kafka-go"
	"gitlab.stackroute.in/JitenP/golang-training-questglobal/models"
)

func main() {
	// make a new reader that consumes from topic-A, partition 0, at offset 42
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:29092"},
		Topic:     "Contact-Created",
		Partition: 0,
		//MinBytes:  10e3, // 10KB
		//MaxBytes:  10e6, // 10MB
	})
	r.SetOffset(6)

	for {
		m, err := r.ReadMessage(context.Background())
		//fmt.Println(err, m)
		if err != nil {
			fmt.Println(err, m)
			//break
		}
		contact := new(models.Contact)
		if err := json.Unmarshal(m.Value, contact); err != nil {
			fmt.Println(contact) // The actual business logic goes
		}
		fmt.Printf("message at offset %d: %s\n", m.Offset, string(m.Value))
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
