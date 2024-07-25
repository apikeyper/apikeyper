package schemas

type CreateKeyResponse struct {
	ApiId string `json:"apiId"`
	KeyId string `json:"keyId"`
	Key   string `json:"key"`
}
