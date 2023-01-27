package handlers

import (
	"context"
	"encoding/json"
	"log"

	"escort-book-user-consumer/config"
	"escort-book-user-consumer/models"
	"escort-book-user-consumer/repositories"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type UserHandler struct {
	DictumRepository         repositories.IDictumRepository
	StatusCategoryRepository repositories.IStatusCategoryRepository
	UserRepository           repositories.IUserRepository
}

func (h *UserHandler) HandleEvent(ctx context.Context, message *kafka.Message) {
	var user models.User
	value := message.Value

	_ = json.Unmarshal(value, &user)

	if err := h.UserRepository.Create(ctx, user); err != nil {
		log.Println("ERROR CREATING THE USER: ", user.Email, err.Error())
		return
	}

	fromUser, err := h.UserRepository.GetByField(ctx, "email", config.InitializeDictum().FromUser)

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

	if err = h.DictumRepository.Create(ctx, newDictum); err != nil {
		log.Println("ERROR CREATING THE DICTUM FOR THE USER: ", user.Email, err.Error())
	}
}
