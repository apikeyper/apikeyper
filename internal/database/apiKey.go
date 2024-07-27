package database

func (s *service) CreateApiKey(apiKey *ApiKey) string {
	s.db.Create(apiKey)
	return apiKey.KeyId
}
