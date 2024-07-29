package server

import "github.com/google/uuid"

type CreateRootKeyResponse struct {
	RootKey string `json:"rootKey"`
}

type CreateApiResponse struct {
	ApiId uuid.UUID `json:"apiId"`
}

type CreateKeyResponse struct {
	ApiId uuid.UUID `json:"apiId"`
	KeyId uuid.UUID `json:"keyId"`
	Key   string    `json:"key"`
}
