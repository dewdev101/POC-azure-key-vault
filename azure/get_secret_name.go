package azure

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *OAuthHandler) GetVualtSecretNameHandler() gin.HandlerFunc {
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

		azure_secret_uri := h.Config.AzureSecretURI
		url := fmt.Sprintf("%s?api-version=2016-10-01", azure_secret_uri)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("error creating request: %v", err)})
			return
		}

		req.Header.Add("Authorization", "Bearer "+token)
		// req.Header.Add("Content-Type", "application/json")

		res, err := h.GetVualtSecretName(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": res})
	}
}
func (h *OAuthHandler) GetVualtSecretName(req *http.Request) (GetVualtSecretNameResponse, error) {

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return GetVualtSecretNameResponse{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return GetVualtSecretNameResponse{}, err
	}

	var response GetVualtSecretNameResponse

	// Unmarshal the JSON into the struct
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
		return GetVualtSecretNameResponse{}, err
	}

	return response, nil
}
