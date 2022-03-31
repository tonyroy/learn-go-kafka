package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	ctx := context.Background()
	l := log.New(os.Stdout, "kafka writer: ", 0)
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic: "purchases",
		Logger: l,
	})

	i := 0
	for {
		err := w.WriteMessages(ctx, kafka.Message{
			Key: []byte(strconv.Itoa(i)),
			Value: []byte("this is message " + strconv.Itoa(i)),
		})
		if err != nil {
			panic("could not write message, " + err.Error())
		}
		fmt.Println("writes: ", i)
		i++
		time.Sleep(time.Second)
	}
	// topic := "purchases"
	
}
