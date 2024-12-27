package database

import (
	"time"

	"github.com/google/uuid"
)

type ApiKeyUsageCount struct {
	IntervalStart time.Time `json:"intervalStart"`
	Success       int       `json:"success"`
	Failed        int       `json:"failed"`
	Revoked       int       `json:"revoked"`
	RateLimited   int       `json:"rateLimited"`
}

type ActivityRecord struct {
	ApiKeyId  uuid.UUID `json:"keyId" gorm:"column:api_key_id"`
	ApiId     uuid.UUID `json:"apiId" gorm:"column:api_id"`
	KeyName   *string   `json:"keyName" gorm:"column:name"` // Added this field
	Usage     string    `json:"usage"`
	Timestamp time.Time `json:"timestamp" gorm:"column:created_at"`
}
