package database

import (
	"fmt"
	"log/slog"
)

func (s *service) CreateUser(user *User) (string, error) {
	result := s.db.Create(user)
	if result.Error != nil {
		slog.Error(fmt.Sprintf("Failed to create user: %v", result.Error))
		return "", result.Error
	}
	slog.Info(fmt.Sprintf("Created user: %v", user.ID))
	return user.GithubId, nil
}

func (s *service) FetchUserByGithubId(githubId string) (*User, error) {
	var user User
	result := s.db.Where("github_id = ?", githubId).Preload("Workspaces").First(&user)
	if result.Error != nil {
		slog.Error(fmt.Sprintf("Failed to fetch user with github id: %s. Error: %v", githubId, result.Error))
		return nil, result.Error
	}
	slog.Info(fmt.Sprintf("Fetched user: %v", user.ID))
	return &user, nil
}
