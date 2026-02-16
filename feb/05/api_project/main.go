package main

import (
	"fmt"
	"log"
	"net"

	"github.com/gabriel-a-costa/LearningGolang/feb/05/api_project/handler"
	"github.com/gabriel-a-costa/LearningGolang/feb/05/api_project/middleware"
	"github.com/gabriel-a-costa/LearningGolang/feb/05/api_project/pb"
	"github.com/gabriel-a-costa/LearningGolang/feb/05/api_project/service"
	"google.golang.org/grpc"
)

func main() {
	//Start service
	userService := service.NewUserService()

	//Set port to listener sever
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor gRPC: %v", err)
	}

	//Config intercept to server gRPC
	serverOptions := []grpc.ServerOption{
		grpc.UnaryInterceptor(
			middleware.UnaryLoggingInterceptor,	
		),
	}

	//Create server gRPC
	//grpcServer := grpc.NewServer()
	
	//Create server gRPC with Middleware
	grpcServer := grpc.NewServer(serverOptions...)

	//Create Handler to abilit methods of server
	userServiceServer := handler.NewUserServiceServer(userService)

	//Register service in server
	pb.RegisterUserServiceServer(
		grpcServer,
		userServiceServer,
	)
	fmt.Println("Server Runner in Port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Erro ao execultar o servidor gRPC: %v", err)
	}
}
