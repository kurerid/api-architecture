package main

import (
	"api-architecture/initialize"
	"api-architecture/pkg/handler"
	"api-architecture/pkg/repository"
	"api-architecture/pkg/service"
	"api-architecture/server"
	"context"
	"github.com/execaus/exloggo"
	"github.com/gin-gonic/gin"
)

func main() {
	exloggo.Info("start server")
	var serverInstance server.Server

	config, err := initialize.Config("configs/config.json")
	if err != nil {
		exloggo.Fatal(err.Error())
		return
	}
	database, err := repository.NewPostgresConnection(config)
	if err != nil {
		exloggo.Fatal(err.Error())
		return
	}

	repos := repository.NewRepository(database)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	ginEngine := handlers.InitRoutes()
	runServer(&serverInstance, ginEngine, "8080")

	serverInstance.Shutdown(database, context.Background())
}

func runServer(server *server.Server, ginEngine *gin.Engine, port string) {
	if err := server.Run(port, ginEngine); err != nil {
		if err.Error() != "http: Server closed" {
			exloggo.Fatalf("error occurred while running http server: %s", nil, err.Error())
		}
	}
}
