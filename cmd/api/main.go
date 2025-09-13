package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"reit-real-estate/config"
	"reit-real-estate/internal/api"
	propertyRepository "reit-real-estate/internal/repository/properties"
	tokenRepository "reit-real-estate/internal/repository/tokens"
	userRepository "reit-real-estate/internal/repository/users"
	userTokenRepository "reit-real-estate/internal/repository/usersTokens"
	walletRepository "reit-real-estate/internal/repository/wallets"
	"reit-real-estate/internal/service"
	"reit-real-estate/pkg/adapter/internal/pkg/reit"
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

	solanaRPC, err := reit.New(reit.Options{
		RPCURL:    cfg.SolanaConfig.RpcURL,
		ProgramID: cfg.SolanaConfig.ProgramID,
	})
	if err != nil {
		log.Println(err)
		return
	}

	userRepo := userRepository.NewRepository(db)
	walletRepo := walletRepository.NewRepository(db)
	propertyRepo := propertyRepository.NewRepository(db)
	tokenRepo := tokenRepository.NewRepository(db)
	userTokenRepo := userTokenRepository.NewRepository(db)
	service := service.NewService(userRepo, walletRepo, propertyRepo, tokenRepo, userTokenRepo, solanaRPC)
	service.WithReit(cfg.SolanaConfig.RpcURL, cfg.SolanaConfig.TokenAddress, cfg.SolanaConfig.ReitMint)

	controller := api.NewController(service)
	server := gin.Default()
	controller.Routes(server)

	err = server.Run(fmt.Sprintf("%s:%d", cfg.AppConfig.Host, cfg.AppConfig.Port))
	if err != nil {
		log.Println(err)
		return
	}
}
