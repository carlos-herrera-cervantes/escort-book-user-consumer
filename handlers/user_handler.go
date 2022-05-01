package handlers

import (
	"context"
	"encoding/json"
	"escort-book-user-consumer/models"
	"escort-book-user-consumer/repositories"
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type UserHandler struct {
	DictumRepository         repositories.IDictumRepository
	StatusCategoryRepository repositories.IStatusCategoryRepository
	UserRepository           repositories.IUserRepository
}

func (h *UserHandler) ProcessMessage(ctx context.Context, message *kafka.Message) {
	var user models.User
	value := message.Value

	json.Unmarshal(value, &user)
	err := h.UserRepository.Create(ctx, user)

	if err != nil {
		log.Println("ERROR CREATING THE USER: ", user.Email, err.Error())
		return
	}

	fromUser, err := h.UserRepository.GetByField(ctx, "email", os.Getenv("FROM_USER"))

	if err != nil {
		log.Println("ERROR GETTING THE USER CREATOR: ", err.Error())
		return
	}

	category, err := h.StatusCategoryRepository.GetByField(ctx, "name", "Active")

	if err != nil {
		log.Println("ERROR GETTING THE STATUS CATEGORY: ", err.Error())
		return
	}

	newDictum := models.Dictum{
		UserId:           user.UserId,
		FromUser:         fromUser.UserId,
		StatusCategoryId: category.Id,
		Comment:          "The user is created by the user consumer component",
	}

	err = h.DictumRepository.Create(ctx, newDictum)

	if err != nil {
		log.Println("ERROR CREATING THE DICTUM FOR THE USER: ", user.Email, err.Error())
	}
}
