package repo

import "context"

//go:generate mockgen -source=repo.go -destination=mocks/mock.go
type Shortener interface {
	SaveURLs(ctx context.Context, shortURL, originalURL string) (string, error)
	GetShortURL(ctx context.Context, originalURL string) (string, error)
	GetOriginalURL(ctx context.Context, shortURL string) (string, error)
}

type Repositories struct {
	Shortener
}

func NewRepositories(shortener Shortener) *Repositories {
	return &Repositories{
		Shortener: shortener,
	}
}
