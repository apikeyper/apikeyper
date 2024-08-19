package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// type User struct {
// 	gorm.Model
// 	ID        uint      `json:"userid" gorm:"primaryKey"`
// 	Email     string    `json:"email"`
// 	Role      string    `json:"role"`
// 	FirstName string    `json:"firstName"`
// 	LastName  string    `json:"lastName"`
// 	CreatedAt time.Time `json:"createdAt"`
// 	UpdatedAt time.Time `json:"updatedAt"`
// }

type Workspace struct {
	gorm.Model
	ID            uuid.UUID `json:"workspaceId" gorm:"primaryKey;type:uuid;default:(gen_random_uuid())"`
	WorkspaceName string    `json:"workspaceName"`
	Apis          []Api
	RootKeys      []RootKey
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

type RootKey struct {
	gorm.Model
	ID            uuid.UUID `json:"rootKeyId" gorm:"primaryKey;type:uuid;default:(gen_random_uuid())"`
	WorkspaceId   uuid.UUID `json:"-"`
	RootHashedKey string    `json:"rootHashedKey"`
	RootKeyName   *string   `json:"rootKeyName"`
	// Permissions   []string  `json:"permissions,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Api struct {
	gorm.Model
	ID          uuid.UUID `json:"apiId" gorm:"primaryKey;type:uuid;default:(gen_random_uuid())"`
	WorkspaceId uuid.UUID `json:"workspaceId"`
	ApiKeys     []ApiKey
	ApiName     string    `json:"apiName"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type ApiKey struct {
	gorm.Model
	ID        uuid.UUID    `json:"apiKeyId" gorm:"primaryKey;type:uuid;default:(gen_random_uuid())"`
	ApiId     uuid.UUID    `json:"-"`
	HashedKey string       `json:"-"` // Store hashed key securely
	Name      *string      `json:"name,omitempty"`
	Prefix    *string      `json:"prefix,omitempty"`
	Status    string       `json:"status" gorm:"default:active"` // active, revoked, expired
	ExpiresAt sql.NullTime `json:"expiresAt"`
	// Roles       []string  `json:"roles,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ApiKeyActivity struct {
	gorm.Model
	ID        uuid.UUID `json:"apiKeyUsageId" gorm:"primaryKey;type:uuid;default:(gen_random_uuid())"`
	ApiKeyId  uuid.UUID `json:"apiKeyId"`
	Usage     string    `json:"usage"` // success, exceeded, rate_limited, revoked
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
