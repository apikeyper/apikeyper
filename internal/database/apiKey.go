package database

import (
	"log"

	"github.com/google/uuid"
)

func (s *service) CreateApiKey(apiKey *ApiKey) uuid.UUID {
	s.db.Create(apiKey)
	return apiKey.ID
}

func (s *service) CreateApi(api *Api) uuid.UUID {
	s.db.Create(api)
	return api.ID
}

func (s *service) FetchApi(apiId uuid.UUID) (*Api, error) {
	var api *Api
	err := s.db.Where("id = ?", apiId).First(&api)

	if err.Error != nil {
		log.Printf("Error: %v", err.Error)
		return nil, err.Error
	}

	log.Printf("Api: %v", api.ID)
	return api, nil
}
