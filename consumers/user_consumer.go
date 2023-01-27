package consumers

import (
	"context"
	"log"

	"escort-book-user-consumer/config"
	"escort-book-user-consumer/handlers"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type UserConsumer struct {
	EventHandler handlers.IEventHandler
}

func (c UserConsumer) StartConsumer() {
	consumer, _ := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": config.InitializeKafka().BootstrapServers,
		"group.id":          config.InitializeKafka().GroupId,
		"auto.offset.reset": "smallest",
	})
	_ = consumer.Subscribe(config.InitializeKafka().Topics.EscortCreated, nil)
	run := true

	for run {
		ev := consumer.Poll(0)

		switch e := ev.(type) {
		case *kafka.Message:
			c.EventHandler.HandleEvent(context.Background(), e)
			log.Println("PROCESSED MESSAGE")
		case kafka.PartitionEOF:
			log.Println("REACHED: ", e)
		case kafka.Error:
			log.Println("KAFKA ERROR: ", e)
			run = false
		default:
		}
	}

	consumer.Close()
}
