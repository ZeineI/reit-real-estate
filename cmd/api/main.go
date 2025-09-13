package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"reit-real-estate/config"
	"reit-real-estate/internal/api"
	userRepository "reit-real-estate/internal/repository/users"
	walletRepository "reit-real-estate/internal/repository/wallets"
	"reit-real-estate/internal/service"
	"reit-real-estate/pkg/postgres"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Println(err)
		return
	}

	db, err := postgres.NewConnection(&postgres.Config{
		Host:               cfg.DatabaseConfig.Host,
		Port:               cfg.DatabaseConfig.Port,
		User:               cfg.DatabaseConfig.User,
		Password:           cfg.DatabaseConfig.Password,
		DatabaseName:       cfg.DatabaseConfig.DatabaseName,
		MaxOpenConnections: cfg.DatabaseConfig.MaxOpenConnections,
		MaxIdleConnections: cfg.DatabaseConfig.MaxIdleConnections,
	})
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	userRepo := userRepository.NewRepository(db)
	walletRepo := walletRepository.NewRepository(db)
	userService := service.NewUserService(userRepo, walletRepo)

	controller := api.NewController(userService)
	server := gin.Default()
	controller.Routes(server)

	err = server.Run(fmt.Sprintf("%s:%s", cfg.AppHost, cfg.AppPort))
	if err != nil {
		log.Println(err)
		return
	}
}
