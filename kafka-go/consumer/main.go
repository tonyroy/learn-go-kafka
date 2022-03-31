package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func main() {
	ctx := context.Background()
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic: "purchases",
		GroupID: "my-group",
	})
	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("Count not read message " + err.Error())
		}
		fmt.Println("Received: ", string(msg.Value))

	}
}