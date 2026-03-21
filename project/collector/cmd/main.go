package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	pb "github.com/Danil-Zaripov/golang-course/project/api/v1"
	"github.com/Danil-Zaripov/golang-course/project/collector/internal/adapters"
	"github.com/Danil-Zaripov/golang-course/project/collector/internal/handlers"
	"github.com/Danil-Zaripov/golang-course/project/collector/internal/usecases"
	"google.golang.org/grpc"
)

const port int = 50051

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	client := &http.Client{}
	adapter := adapters.NewGHFetcher(client)
	usecase := usecases.NewRepoUsecase(adapter)
	handler := handlers.NewRepoHandler(usecase)

	pb.RegisterCollectorServer(s, handler)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Printf("failed to serve: %v", err)
	}
}
