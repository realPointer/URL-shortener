package service

import (
	"context"

	"github.com/realPointer/url-shortener/internal/repo"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go
type Shortener interface {
	GetShortURL(ctx context.Context, originalURL string) (string, error)
	GetOriginalURL(ctx context.Context, shortURL string) (string, error)
}

type Services struct {
	Shortener
}

type ServicesDependencies struct {
	Repo *repo.Repositories
}

func NewServices(deps ServicesDependencies) *Services {
	return &Services{
		Shortener: NewShortenerService(deps.Repo.Shortener),
	}
}
