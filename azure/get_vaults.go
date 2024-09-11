package azure

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewGetListofVaults() *OAuthHandler {
	return &OAuthHandler{}
}
func (h *OAuthHandler) GetListofVaultsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the token from the context set by the middleware
		token, exists := c.Get("token")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is missing"})
			return
		}

		// Call the method to get the list of vaults using the token
		res, err := h.GetListofVaults(token.(string))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": res})
	}
}

func (h *OAuthHandler) GetListofVaults(token string) (GetListofVaultsResponse, error) {

	subscriptionId := h.Config.SubscriptionId
	url := "https://management.azure.com/subscriptions/" + subscriptionId + "/resources?%24filter=resourceType%20eq%20%27Microsoft.KeyVault%2Fvaults%27&api-version=2024-03-01"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return GetListofVaultsResponse{}, nil
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return GetListofVaultsResponse{}, nil
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return GetListofVaultsResponse{}, nil
	}

	rest := GetListofVaultsResponse{}
	err = json.Unmarshal(body, &rest)
	if err != nil {
		fmt.Println(err)
		return GetListofVaultsResponse{}, err
	}
	return rest, err
}
