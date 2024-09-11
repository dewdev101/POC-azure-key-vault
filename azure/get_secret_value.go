package azure

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *OAuthHandler) GetSecretValueHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization token is required"})
			return
		}

		// Remove "Bearer " prefix if present
		if len(token) > 7 && token[:7] == "Bearer " {
			token = token[7:]
		}

		res, err := h.GetSecretValue(token)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": res})
	}
}

func (h *OAuthHandler) GetSecretValue(token string) (AzureKeyVaultValue, error) {

	azureSecretURI := h.Config.AzureSecretURI

	// Fetch the private key from Azure Key Vault
	url := fmt.Sprintf("%s/privateKey?api-version=2016-10-01", azureSecretURI)
	fmt.Printf("url:%s\n", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return AzureKeyVaultValue{}, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return AzureKeyVaultValue{}, fmt.Errorf("error making request: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return AzureKeyVaultValue{}, fmt.Errorf("error reading response body: %v", err)
	}

	var secret SecretValue
	if err := json.Unmarshal(body, &secret); err != nil {
		return AzureKeyVaultValue{}, fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	// Fetch another secret (if needed)
	url2 := fmt.Sprintf("%s/secret?api-version=2016-10-01", azureSecretURI)
	req, err = http.NewRequest("GET", url2, nil)
	if err != nil {
		return AzureKeyVaultValue{}, fmt.Errorf("error creating second request: %v", err)
	}
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")

	res, err = client.Do(req)
	if err != nil {
		return AzureKeyVaultValue{}, fmt.Errorf("error making second request: %v", err)
	}
	defer res.Body.Close()

	body, err = io.ReadAll(res.Body)
	if err != nil {
		return AzureKeyVaultValue{}, fmt.Errorf("error reading second response body: %v", err)
	}

	var secret2 SecretValue
	if err := json.Unmarshal(body, &secret2); err != nil {
		return AzureKeyVaultValue{}, fmt.Errorf("error unmarshalling second JSON: %v", err)
	}

	response := AzureKeyVaultValue{
		PrivateKey: secret.Value,
		Secret:     secret2.Value,
	}
	// encrypt and save value to keyFile
	if response.PrivateKey != "" || response.Secret != "" {
		Encrypt(secret.Value, secret2.Value)
	}

	return response, nil

}
