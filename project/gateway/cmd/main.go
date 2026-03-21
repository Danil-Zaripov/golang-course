package main

import (
	"fmt"
	"log"
	"net/http"

	pb "github.com/Danil-Zaripov/golang-course/project/api/v1"
	"github.com/Danil-Zaripov/golang-course/project/gateway/internal/adapters"
	"github.com/Danil-Zaripov/golang-course/project/gateway/internal/handlers"
	"github.com/Danil-Zaripov/golang-course/project/gateway/internal/usecases"
	httpSwagger "github.com/swaggo/http-swagger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	_ "github.com/Danil-Zaripov/golang-course/project/gateway/cmd/docs"
)

const grpcport int = 50051
const port int = 50031

//	@title			GitHub Repo fetcher API
//	@version		1.0
//	@description	Gets some repository data from github api

//	@license.name	MIT

// @host	localhost:50031
func main() {
	conn, err := grpc.NewClient(fmt.Sprintf(":%d", grpcport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCollectorClient(conn)

	adapter := adapters.NewCollectorAdapter(c)
	usecase := usecases.NewRepoUsecase(adapter)
	handler := handlers.NewRepoHandler(usecase)

	http.HandleFunc("/api/repo/", handler.HandleRepo)
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	fmt.Printf("Starting listening on :%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
