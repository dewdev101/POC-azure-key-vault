package azure

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func (h *OAuthHandler) GetOAuthTokenSecretHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		token, err := h.GetOAuthSecretToken()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

func (h *OAuthHandler) GetOAuthSecretToken() (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	tenantId := h.Config.TenantId

	url := "https://login.microsoftonline.com/" + tenantId + "/oauth2/token"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("grant_type", "client_credentials")
	_ = writer.WriteField("client_id", h.Config.ClientId)
	_ = writer.WriteField("client_secret", h.Config.ClientSecret)
	_ = writer.WriteField("resource", "https://vault.azure.net")
	err = writer.Close()
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Cookie", "fpc=AsNpSSQrk0NIvTL8ERnVp6RmGgaQAQAAACoict4OAAAA; stsservicecookie=estsfd; x-ms-gateway-slice=estsfd")

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	rest := GetOAuthTokenResponse{}
	err = json.Unmarshal(body, &rest)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return rest.AccessToken, nil
}
