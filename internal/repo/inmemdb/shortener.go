package inmemdb

import (
	"context"
	"fmt"
	"sync"

	"github.com/realPointer/url-shortener/internal/repo"
)

var _ repo.Shortener = &ShortenerRepo{}

type ShortenerRepo struct {
	originalStorage map[string]string
	shortStorage    map[string]string
	mutexOriginal   sync.RWMutex
	mutexShort      sync.RWMutex
}

func NewShortenerRepo() *ShortenerRepo {
	return &ShortenerRepo{
		originalStorage: make(map[string]string),
		shortStorage:    make(map[string]string),
	}
}

func (s *ShortenerRepo) SaveURLs(ctx context.Context, shortURL, originalURL string) (string, error) {
	s.mutexOriginal.Lock()
	defer s.mutexOriginal.Unlock()

	s.mutexShort.Lock()
	defer s.mutexShort.Unlock()

	s.originalStorage[originalURL] = shortURL
	s.shortStorage[shortURL] = originalURL

	return s.originalStorage[originalURL], nil
}

func (s *ShortenerRepo) GetShortURL(ctx context.Context, originalURL string) (string, error) {
	s.mutexOriginal.RLock()
	defer s.mutexOriginal.RUnlock()

	shortURL, ok := s.originalStorage[originalURL]
	if ok {
		return shortURL, nil
	}

	return "", fmt.Errorf("not found")
}

func (s *ShortenerRepo) GetOriginalURL(ctx context.Context, shortURL string) (string, error) {
	s.mutexShort.RLock()
	defer s.mutexShort.RUnlock()

	originalURL, ok := s.shortStorage[shortURL]
	if ok {
		return originalURL, nil
	}

	return "", fmt.Errorf("not found")
}
