package main

import (
	"io"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"
)

// pack transferFrom function
func PackTransferFromData(from, to common.Address, value *big.Int) ([]byte, error) {
	contractAbi, err := LoadABI("./token.json")
	if err != nil {
		return nil, err
	}

	parsedABI, err := abi.JSON(strings.NewReader(contractAbi))
	if err != nil {
		return nil, err
	}

	// Pack the method name and arguments into call data
	callData, err := parsedABI.Pack("transferFrom", from, to, value)
	if err != nil {
		return nil, err
	}

	return callData, nil
}
func LoadABI(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func Encrypt(secret string) {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Retrieve private key and password from environment variables
	privateKeyHex := os.Getenv("METAMASK_PRIVATE_KEY")
	password := os.Getenv("PASSWORD")

	if privateKeyHex == "" || password == "" {
		log.Fatal("METAMASK_PRIVATE_KEY or PASSWORD environment variable is not set")
	}

	// Convert private key hex string to *ecdsa.PrivateKey
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("Failed to convert private key hex to ECDSA: %v", err)
	}

	// Create a temporary directory for the keystore
	keystoreDir := "ether-signer"
	os.MkdirAll(keystoreDir, 0700)
	ks := keystore.NewKeyStore(keystoreDir, keystore.StandardScryptN, keystore.StandardScryptP)

	// Generate the address from the private key
	address := crypto.PubkeyToAddress(privateKey.PublicKey).Hex()

	// Check if the account already exists in the keystore
	var account *accounts.Account
	for _, acc := range ks.Accounts() {
		if acc.Address.Hex() == address {
			account = &acc
			break
		}
	}

	if account == nil {
		// Import the private key into the keystore
		acc, err := ks.ImportECDSA(privateKey, password)
		if err != nil {
			log.Fatalf("Failed to import ECDSA key: %v", err)
		}
		account = &acc
	}

	// Export the keystore JSON
	keyJSON, err := ks.Export(*account, password, password)
	if err != nil {
		log.Fatalf("Failed to export keystore: %v", err)
	}

	// Save the keystore JSON to a file named "keyfile" in the keystore directory
	keystoreFilePath := filepath.Join(keystoreDir, "keyfile")
	err = os.WriteFile(keystoreFilePath, keyJSON, 0600)
	if err != nil {
		log.Fatalf("Failed to save keystore to file: %v", err)
	}

	// Save the secret (password) to a file named "secret" in the "ether-signer2" directory
	secretFilePath := filepath.Join(keystoreDir, "passwordfile")
	err = os.WriteFile(secretFilePath, []byte(secret), 0600)
	if err != nil {
		log.Fatalf("Failed to save secret to file: %v", err)
	}

	exceptions := []string{"keyfile", "passwordfile"}

	// Clean the directory
	err = CleanDirectory(keystoreDir, exceptions)
	if err != nil {
		log.Fatalf("Error cleaning directory: %v", err)
	}

}
func CleanDirectory(dir string, exceptions []string) error {
	// Convert exceptions to a map for quick lookup
	exceptionMap := make(map[string]bool)
	for _, name := range exceptions {
		exceptionMap[name] = true
	}

	// Walk through the directory
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories and exception files
		if !info.IsDir() && !exceptionMap[info.Name()] {
			// Delete the file
			if err := os.Remove(path); err != nil {
				return err
			}
		}
		return nil
	})

	return err
}
