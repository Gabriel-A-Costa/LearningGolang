package service

import (
	"errors"

	"github.com/gabriel-a-costa/LearningGolang/feb/04/api_project/model"
)

var users = []model.User{
	{ID: "1", Name: "Alice", Age: 30},
	{ID: "2", Name: "Bod", Age: 25},
}

func GetAllUsers() []model.User {
	return users
}

func GetUserById(id string) (*model.User, error) {
	for _, user := range users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("Usuário não encontrado")
}
