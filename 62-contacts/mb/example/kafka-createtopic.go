package main

import (
	"context"

	kafka "github.com/segmentio/kafka-go"
)

func main() {
	// to create topics when auto.create.topics.enable='true'
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:29092", "Contact-Created", 0)
	if err != nil {
		panic(err.Error())
	}

	conn.Close()
}
