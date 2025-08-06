package tools

import (
	"context"
	"errors"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	//brokers := "b-2.nsmspdev.04141j.c3.kafka.us-west-2.amazonaws.com:9096,b-1.nsmspdev.04141j.c3.kafka.us-west-2.amazonaws.com:9096"
	conn, err := kafka.Dial("tcp", "b-2.nsmspdev.04141j.c3.kafka.us-west-2.amazonaws.com:9096")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	partitions, err := conn.ReadPartitions()
	if err != nil {
		panic(err.Error())
	}

	m := map[string]struct{}{}

	for _, p := range partitions {
		m[p.Topic] = struct{}{}
	}
	for k := range m {
		fmt.Println(k)
	}
}

func TestName2(t *testing.T) {
	// Make a writer that publishes messages to topic-A.
	// The topic will be created if it is missing.
	w := &kafka.Writer{
		Addr: kafka.TCP(
			"b-2.nsmspdev.04141j.c3.kafka.us-west-2.amazonaws.com:9096",
			"b-1.nsmspdev.04141j.c3.kafka.us-west-2.amazonaws.com:9096",
		),
		Topic:                  "topic-A",
		AllowAutoTopicCreation: true,
	}

	messages := []kafka.Message{
		{
			Key:   []byte("Key-A"),
			Value: []byte("Hello World!"),
		},
		{
			Key:   []byte("Key-B"),
			Value: []byte("One!"),
		},
		{
			Key:   []byte("Key-C"),
			Value: []byte("Two!"),
		},
	}

	var err error
	const retries = 3
	for i := 0; i < retries; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// attempt to create topic prior to publishing the message
		err = w.WriteMessages(ctx, messages...)
		if errors.Is(err, kafka.LeaderNotAvailable) || errors.Is(err, context.DeadlineExceeded) {
			time.Sleep(time.Millisecond * 250)
			continue
		}

		if err != nil {
			log.Fatalf("unexpected error %v", err)
		}
		break
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
