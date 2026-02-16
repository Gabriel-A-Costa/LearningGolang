package handler

import (
	"context"
	"fmt"

	pb "github.com/gabriel-a-costa/LearningGolang/feb/05/api_project/pb"
	"github.com/gabriel-a-costa/LearningGolang/feb/05/api_project/service"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
	service *service.UserService
}

// This is a constructor
func NewUserServiceServer(service *service.UserService) *UserServiceServer {
	return &UserServiceServer{service: service}
}

func (s *UserServiceServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	id, err := s.service.CreateUser(req.Name, int(req.Age))
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{
		Id:      id,
		Message: "Usuário criado com sucesso!",
	}, nil
}

func (s *UserServiceServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := s.service.GetUserByID(req.Id)
	if err != nil {
		return nil, fmt.Errorf("Usuário não encontrado!")
	}

	return &pb.GetUserResponse{
		Id:   user.ID,
		Name: user.Name,
		Age:  int32(user.Age),
	}, nil
}
