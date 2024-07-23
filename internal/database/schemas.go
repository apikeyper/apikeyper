package database

import (
	"time"

	"gorm.io/gorm"
)

type ApiKey struct {
	gorm.Model
	KeyId     string
	KeyHash   string
	CreatedAt time.Time
}
