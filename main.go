package main

import (
	"escort-book-user-consumer/consumers"
	"escort-book-user-consumer/db"
	"escort-book-user-consumer/handlers"
	"escort-book-user-consumer/repositories"
)

func main() {
	userHandler := &handlers.UserHandler{
		DictumRepository: &repositories.DictumRepository{
			Data: db.NewPostgresClient(),
		},
		StatusCategoryRepository: &repositories.StatusCategoryRepository{
			Data: db.NewPostgresClient(),
		},
		UserRepository: &repositories.UserRepository{
			Data: db.NewPostgresClient(),
		},
	}
	userConsumer := consumers.UserConsumer{
		EventHandler: userHandler,
	}

	userConsumer.StartConsumer()
}
