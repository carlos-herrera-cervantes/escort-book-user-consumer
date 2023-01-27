package config

import "os"

type kafka struct {
	BootstrapServers string
	GroupId          string
	Topics           topic
}

type topic struct {
	EscortCreated     string
	UserActiveAccount string
}

var singleKafka *kafka

func InitializeKafka() *kafka {
	if singleKafka != nil {
		return singleKafka
	}

	lock.Lock()
	defer lock.Unlock()

	singleKafka = &kafka{
		BootstrapServers: os.Getenv("KAFKA_SERVERS"),
		GroupId:          os.Getenv("KAFKA_GROUP_ID"),
		Topics: topic{
			EscortCreated:     "user-created",
			UserActiveAccount: "",
		},
	}

	return singleKafka
}
