package service

import (
	"context"

	"github.com/realPointer/url-shortener/internal/repo"
	"github.com/realPointer/url-shortener/pkg/urlgenerator"
)

type ShortenerService struct {
	shortenerRepo repo.Shortener
}

func NewShortenerService(shortenerRepo repo.Shortener) *ShortenerService {
	return &ShortenerService{
		shortenerRepo: shortenerRepo,
	}
}

func (s *ShortenerService) GetShortURL(ctx context.Context, originalURL string) (string, error) {
	shortURL, err := s.shortenerRepo.GetShortURL(ctx, originalURL)
	var shortenedURL string
	if err != nil {

		for {
			shortenedURL = urlgenerator.GenerateShortURL(originalURL)

			// Check if the generated short URL already exists in the repository
			_, err := s.shortenerRepo.GetOriginalURL(ctx, shortenedURL)
			if err != nil {
				shortURL, _ = s.shortenerRepo.SaveURLs(ctx, shortenedURL, originalURL)
				break
			}
		}
	}

	return shortURL, nil
}

func (s *ShortenerService) GetOriginalURL(ctx context.Context, shortURL string) (string, error) {
	return s.shortenerRepo.GetOriginalURL(ctx, shortURL)
}
