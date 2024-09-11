package azure

type AzureKeyVaultValue struct {
	PrivateKey string `json:"privateKey"`
	Secret     string `json:"secret"`
}
type AttributeSecret struct {
	Enabled       bool   `json:"enabled"`
	Created       int64  `json:"created"`
	Updated       int64  `json:"updated"`
	RecoveryLevel string `json:"recoveryLevel"`
}

type SecretValue struct {
	Value      string                 `json:"value"`
	ID         string                 `json:"id"`
	Attributes Attributes             `json:"attributes"`
	Tags       map[string]interface{} `json:"tags"`
}

// Define the structs based on the JSON structure
type Attributes struct {
	Enabled       bool   `json:"enabled"`
	Created       int64  `json:"created"`
	Updated       int64  `json:"updated"`
	RecoveryLevel string `json:"recoveryLevel"`
}

type Secret struct {
	ID         string                 `json:"id"`
	Attributes Attributes             `json:"attributes"`
	Tags       map[string]interface{} `json:"tags"`
}

type GetVualtSecretNameResponse struct {
	Value    []Secret `json:"value"`
	NextLink *string  `json:"nextLink"`
}
type GetOAuthTokenResponse struct {
	TokenType    string `json:"token_type" `
	ExpireIn     string `json:"expires_in"`
	ExtExpiresIn string `json:"ext_expires_in"`
	ExpiresOn    string `json:"expires_on"`
	NotBefore    string `json:"not_before"`
	Resource     string `json:"resource"`
	AccessToken  string `json:"access_token"`
}

type GetListofVaultsResponse struct {
	Value []GetListofVaultsResponseDetail `json:"value"`
}
type GetListofVaultsResponseDetail struct {
	Id       string            `json:"id"`
	Name     string            `json:"name"`
	Type     string            `json:"type"`
	Location string            `json:"location"`
	Tags     map[string]string `json:"tags"`
}
