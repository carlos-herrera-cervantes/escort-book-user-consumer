package main

import (
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	consumer, _ := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_SERVERS"),
		"group.id":          os.Getenv("KAFKA_GROUP_ID"),
		"auto.offset.reset": "smallest",
	})
	run := true
	consumer.Subscribe(os.Getenv("KAFKA_USER_TOPIC"), nil)

	for run {
		ev := consumer.Poll(0)

		switch e := ev.(type) {
		case *kafka.Message:
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
