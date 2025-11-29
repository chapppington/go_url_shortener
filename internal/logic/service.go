package logic

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"urlshortener/internal/infrastructure"
)

type Service struct {
	repo *infrastructure.Repository
}

func NewService(repo *infrastructure.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateShortURL(longURL string) (*infrastructure.ShortURL, error) {
	shortCode, err := s.generateShortCode()
	if err != nil {
		return nil, fmt.Errorf("failed to generate short code: %w", err)
	}

	for {
		existing, err := s.repo.FindByShortCode(shortCode)
		if err != nil {
			return nil, fmt.Errorf("failed to check short code: %w", err)
		}
		if existing == nil {
			break
		}
		shortCode, err = s.generateShortCode()
		if err != nil {
			return nil, fmt.Errorf("failed to generate short code: %w", err)
		}
	}

	shortURL := &infrastructure.ShortURL{
		LongUrl:   longURL,
		ShortCode: shortCode,
	}

	if err := s.repo.Create(shortURL); err != nil {
		return nil, fmt.Errorf("failed to create short URL: %w", err)
	}

	return shortURL, nil
}

func (s *Service) GetLongURL(shortCode string) (string, error) {
	shortURL, err := s.repo.FindByShortCode(shortCode)
	if err != nil {
		return "", fmt.Errorf("failed to find short URL: %w", err)
	}
	if shortURL == nil {
		return "", fmt.Errorf("short code not found")
	}
	return shortURL.LongUrl, nil
}

func (s *Service) generateShortCode() (string, error) {
	bytes := make([]byte, 6)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes)[:8], nil
}