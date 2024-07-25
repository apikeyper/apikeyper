package database

import (
	"time"

	"gorm.io/gorm"
)

type ApiKey struct {
	gorm.Model
	KeyId       string `json:"apiKeyId" gorm:"primaryKey"`
	ApiId       string `json:"apiId"`
	WorkspaceId string `json:"workspaceId"`
	HashedKey   string `json:"-"` // Store hashed key securely
	// Name        *string `json:"name,omitempty"`
	// Prefix      *string `json:"prefix,omitempty"`
	// Roles            []string  `json:"roles,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
