package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"poc-azure-key-vault/azure"
	config "poc-azure-key-vault/config"
	"poc-azure-key-vault/middleware"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	config, err := config.Configuration()
	if err != nil {
		log.Fatalf("Error loading .config file: %v", err)
	}

	r := gin.Default()

	oauthHandler := azure.NewOAuthHandler(config)

	r.POST("/token", oauthHandler.GetOAuthTokenHandler())
	r.GET("/vaults", middleware.AuthorizationMiddleware(), oauthHandler.GetListofVaultsHandler())
	r.POST("/token/secret", oauthHandler.GetOAuthTokenSecretHandler())
	r.GET("/vault/secret-name", middleware.AuthorizationMiddleware(), oauthHandler.GetVualtSecretNameHandler())
	r.GET("/secret", middleware.AuthorizationMiddleware(), oauthHandler.GetSecretValueHandler())

	r.Run(":8080")

	// ====== prepare data for send eth

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// ==== send transfer function

	from := common.HexToAddress(os.Getenv("FROM_CONTRACT_ADDRESS"))
	to := common.HexToAddress(os.Getenv("TO_CONTRACT_ADDRESS"))
	value, _ := big.NewInt(0).SetString("10000000000000000000", 10)
	if value == nil {
		fmt.Println("Error creating big.Int")
		return
	}

	// ====== pack method transferFrom to call data to call eth_send_tr

	callData, err := PackTransferFromData(from, to, value)
	if err != nil {
		log.Fatalf("Failed to pack transferFrom data: %v", err)
	}

	log.Printf("Call data: %x", callData)

}
