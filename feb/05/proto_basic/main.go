package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/gabriel-a-costa/LearningGolang/feb/05/proto_basic/pb"
	"google.golang.org/grpc"
)

type GreeterServer struct {
	pb.UnimplementedGreeterServer
}

type Body struct {
	Name         string
	Age          int64
	Is_brazilian bool
}

func (g *GreeterServer) SayHello(ctx context.Context, r *pb.SayHelloRequest) (*pb.SayHellorResponse, error) {
	body := Body{
		Name:         r.Name,
		Age:          r.Idade,
		Is_brazilian: r.IsBrazilian,
	}

	return &pb.SayHellorResponse{
		Message: fmt.Sprintf("Olá, %s! Bem-vindo ao mundo gRPC! sua idade é %d, você é brasileiro? %t", body.Name, body.Age, body.Is_brazilian),
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &GreeterServer{})
	fmt.Println("Servidor gRPC rodando na porta 50051...")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Erro ao execultar o servidor: %v", err)
	}
}
