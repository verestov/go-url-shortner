package main

import (
	_ "context"
	"fmt"
	"log"
	"net"

	shortenerpb "github.com/verestov/go-url-shortner.git/gen" // импорт сгенерированного кода

	"github.com/verestov/go-url-shortner.git/internal/handler"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	// Регистрируем сервис
	// Реализация нашего сервиса через gRPC хендлер
	shortenerpb.RegisterShortenerServiceServer(s, &handler.Server{})

	fmt.Println("Server running on port", port)
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
