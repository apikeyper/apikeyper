package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         string       `json:"userId" gorm:"primaryKey"`
	Username   string       `json:"username"`
	GithubId   string       `json:"githubId" gorm:"uniqueIndex:github_id_unq_idx"`
	Workspaces []*Workspace `gorm:"many2many:user_workspaces;"`
}

type Session struct {
	gorm.Model
	ID        string `json:"sessionId" gorm:"primaryKey"`
	UserId    string `json:"userId"`
	ExpiresAt time.Time
}

type Workspace struct {
	gorm.Model
	ID            uuid.UUID `json:"workspaceId" gorm:"primaryKey;type:uuid;default:(gen_random_uuid())"`
	WorkspaceName string    `json:"workspaceName"`
	Users         []*User   `gorm:"many2many:user_workspaces;"`
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
	ID              uuid.UUID             `json:"apiKeyId" gorm:"primaryKey;type:uuid;default:(gen_random_uuid())"`
	ApiId           uuid.UUID             `json:"-"`
	HashedKey       string                `json:"-"` // Store hashed key securely
	Name            *string               `json:"name,omitempty"`
	Prefix          *string               `json:"prefix,omitempty"`
	Status          string                `json:"status" gorm:"default:active"` // active, revoked, expired
	ExpiresAt       sql.NullTime          `json:"expiresAt"`
	RateLimitConfig ApiKeyRateLimitConfig `gorm:"constraint:OnDelete:CASCADE;"`
	// Roles       []string  `json:"roles,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ApiKeyRateLimitConfig struct {
	gorm.Model
	ID            uuid.UUID `json:"apiKeyRateLimitConfigId" gorm:"primaryKey;type:uuid;default:(gen_random_uuid())"`
	ApiKeyId      uuid.UUID `json:"-"`
	Limit         int       `json:"limit"`
	LimitPeriod   time.Duration
	CounterWindow time.Duration
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

type ApiKeyActivity struct {
	gorm.Model
	ID        uuid.UUID `json:"apiKeyUsageId" gorm:"primaryKey;type:uuid;default:(gen_random_uuid())"`
	ApiKeyId  uuid.UUID `json:"apiKeyId"`
	Usage     string    `json:"usage"` // success, exceeded, rate_limited, revoked
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
