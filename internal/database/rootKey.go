package database

import (
	"log"

	"github.com/google/uuid"
)

func (s *service) CreateRootKey(rootKey *RootKey) (uuid.UUID, error) {
	tx := s.db.Create(rootKey)
	if tx.Error != nil {
		log.Printf("Error: %v", tx.Error)
		return uuid.Nil, tx.Error
	}

	return rootKey.ID, nil
}

func (s *service) FetchRootKey(rootHashedKey string) *RootKey {
	var rootKey RootKey
	if err := s.db.Where("root_hashed_key = ?", rootHashedKey).First(&rootKey); err.Error != nil {
		log.Printf("Error: %v", err.Error)
		return nil
	}

	log.Printf("RootKey: %v", rootKey.RootHashedKey)
	return &rootKey
}
