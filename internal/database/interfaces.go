package database

import (
	"time"
)

type ApiKeyUsageCount struct {
	IntervalStart time.Time `json:"intervalStart"`
	Success       int       `json:"success"`
	Failed        int       `json:"failed"`
	Revoked       int       `json:"revoked"`
	RateLimited   int       `json:"rateLimited"`
}
