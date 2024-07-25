package database

import "log"

func (s *service) CreateApiKey(apiKey *ApiKey) string {
	log.Print("3")
	s.db.Create(apiKey)
	return apiKey.KeyId
}
