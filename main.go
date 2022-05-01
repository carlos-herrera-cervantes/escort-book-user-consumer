package main

import (
	"context"
	"escort-book-user-consumer/db"
	"escort-book-user-consumer/handlers"
	"escort-book-user-consumer/repositories"
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

	handler := handlers.UserHandler{
		DictumRepository: &repositories.DictumRepository{
			Data: db.New(),
		},
		StatusCategoryRepository: &repositories.StatusCategoryRepository{
			Data: db.New(),
		},
		UserRepository: &repositories.UserRepository{
			Data: db.New(),
		},
	}

	for run {
		ev := consumer.Poll(0)

		switch e := ev.(type) {
		case *kafka.Message:
			handler.ProcessMessage(context.Background(), e)
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
