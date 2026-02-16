package service

import (
	"errors"
	"fmt"

	"github.com/gabriel-a-costa/LearningGolang/feb/04/api_project/model"
)

type UserService struct {
	users map[string]*model.User
}

// This is a constructor
func NewUserService() *UserService {
	return &UserService{
		users: make(map[string]*model.User),
	}
}

func (s *UserService) CreateUser(name string, age int) (string, error) {
	id := fmt.Sprintf("%d", len(s.users)+1)
	s.users[id] = &model.User{
		ID:   id,
		Name: name,
		Age:  age,
	}

	return id, nil
}

func (s *UserService) GetUserByID(id string) (*model.User, error) {
	user, exist := s.users[id]
	if !exist {
		return nil, errors.New("Usuário não encontrado!")
	}
	return user, nil
}
