package handlers

import (
	"context"
	"errors"
	"testing"

	"escort-book-user-consumer/models"
	"escort-book-user-consumer/repositories/mocks"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/golang/mock/gomock"
)

func TestUserHandlerHandleEvent(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDictumRepository := mocks.NewMockIDictumRepository(ctrl)
	mockStatusCategoryRepository := mocks.NewMockIStatusCategoryRepository(ctrl)
	mockUserRepository := mocks.NewMockIUserRepository(ctrl)

	userHandler := UserHandler{
		DictumRepository:         mockDictumRepository,
		StatusCategoryRepository: mockStatusCategoryRepository,
		UserRepository:           mockUserRepository,
	}

	t.Run("Should exit when creating user fails", func(t *testing.T) {
		mockUserRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(errors.New("dummy error")).
			Times(1)
		mockUserRepository.
			EXPECT().
			GetByField(gomock.Any(), gomock.Any(), gomock.Any()).
			Times(0)
		mockStatusCategoryRepository.
			EXPECT().
			GetByField(gomock.Any(), gomock.Any(), gomock.Any()).
			Times(0)
		mockDictumRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Times(0)

		kafkaMessage := kafka.Message{
			Value: []byte(`{
                "_id": "63d299a588a911dfc86370cc",
                "email": "user@example.com"
            }`),
		}
		userHandler.HandleEvent(context.Background(), &kafkaMessage)
	})

	t.Run("Should exit when getting user fails", func(t *testing.T) {
		mockUserRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(nil).
			Times(1)
		mockUserRepository.
			EXPECT().
			GetByField(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(models.User{}, errors.New("dummy error")).
			Times(1)
		mockStatusCategoryRepository.
			EXPECT().
			GetByField(gomock.Any(), gomock.Any(), gomock.Any()).
			Times(0)
		mockDictumRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Times(0)

		kafkaMessage := kafka.Message{
			Value: []byte(`{
                "_id": "63d299a588a911dfc86370cc",
                "email": "user@example.com"
            }`),
		}
		userHandler.HandleEvent(context.Background(), &kafkaMessage)
	})

	t.Run("Should exit when getting status category fails", func(t *testing.T) {
		mockUserRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(nil).
			Times(1)
		mockUserRepository.
			EXPECT().
			GetByField(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(models.User{UserId: "63d29e3e0b68c6beaea6fe4f"}, nil).
			Times(1)
		mockStatusCategoryRepository.
			EXPECT().
			GetByField(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(models.StatusCategory{}, errors.New("dummy error")).
			Times(1)
		mockDictumRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Times(0)

		kafkaMessage := kafka.Message{
			Value: []byte(`{
                "_id": "63d299a588a911dfc86370cc",
                "email": "user@example.com"
            }`),
		}
		userHandler.HandleEvent(context.Background(), &kafkaMessage)
	})

	t.Run("Should log error when creating dictum fails", func(t *testing.T) {
		mockUserRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(nil).
			Times(1)
		mockUserRepository.
			EXPECT().
			GetByField(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(models.User{UserId: "63d29e3e0b68c6beaea6fe4f"}, nil).
			Times(1)
		mockStatusCategoryRepository.
			EXPECT().
			GetByField(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(models.StatusCategory{Id: "63d29ee74efb1183cd08a985"}, nil).
			Times(1)
		mockDictumRepository.
			EXPECT().
			Create(gomock.Any(), gomock.Any()).
			Return(errors.New("dummy error")).
			Times(1)

		kafkaMessage := kafka.Message{
			Value: []byte(`{
                "_id": "63d299a588a911dfc86370cc",
                "email": "user@example.com"
            }`),
		}
		userHandler.HandleEvent(context.Background(), &kafkaMessage)
	})
}
