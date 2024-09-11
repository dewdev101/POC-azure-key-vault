package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AzureSecretURI string
	SubscriptionId string
	ClientId       string
	ClientSecret   string
	TenantId       string
}

func Configuration() (*Config, error) {

	config := Config{}
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
		return nil, err
	}

	config.AzureSecretURI = os.Getenv("AZURE_SECRET_URI")
	if config.AzureSecretURI == "" {
		return nil, fmt.Errorf("AZURE_SECRET_URI not set")
	}
	config.SubscriptionId = os.Getenv("SUBSCRIPTION_ID")
	if config.SubscriptionId == "" {
		return nil, fmt.Errorf("SubscriptionId not set")
	}
	config.ClientId = os.Getenv("CLIENT_ID")
	if config.ClientId == "" {
		return nil, fmt.Errorf("ClientId not set")
	}
	config.ClientSecret = os.Getenv("CLIENT_SECRET")
	if config.ClientSecret == "" {
		return nil, fmt.Errorf("CLIENT_SECRET not set")
	}
	config.TenantId = os.Getenv("TENANT_ID")
	if config.TenantId == "" {
		return nil, fmt.Errorf("TENANT_ID not set")
	}
	return &config, nil
}
