package postgresrepo

import (
	"context"
	"fmt"

	"github.com/realPointer/url-shortener/internal/repo"
	"github.com/realPointer/url-shortener/pkg/postgres"
)

var _ repo.Shortener = &ShortenerRepo{}

type ShortenerRepo struct {
	*postgres.Postgres
}

func NewShortenerRepo(pg *postgres.Postgres) *ShortenerRepo {
	return &ShortenerRepo{
		Postgres: pg,
	}
}

func (s *ShortenerRepo) SaveURLs(ctx context.Context, shortURL, originalURL string) (string, error) {
	sql, args, _ := s.Builder.
		Insert("links").
		Columns("shortURL", "originalURL").
		Values(shortURL, originalURL).
		Suffix("RETURNING links.shortURL").
		ToSql()

	var res string
	row := s.Pool.QueryRow(ctx, sql, args...)
	err := row.Scan(&res)
	if err != nil {
		return "", fmt.Errorf("ShortenerRepo - SaveShortURL - row.Scan: %w", err)
	}

	return res, nil
}

func (s *ShortenerRepo) GetShortURL(ctx context.Context, originalURL string) (string, error) {
	sql, args, _ := s.Builder.
		Select("shortURL").
		From("links").
		Where("originalURL = ?", originalURL).
		ToSql()

	var shortURL string
	row := s.Pool.QueryRow(ctx, sql, args...)
	err := row.Scan(&shortURL)
	if err != nil {
		return "", fmt.Errorf("ShortenerRepo - getShortURL - row.Scan: %w", err)
	}

	return shortURL, nil
}

func (s *ShortenerRepo) GetOriginalURL(ctx context.Context, shortURL string) (string, error) {
	sql, args, _ := s.Builder.
		Select("originalURL").
		From("links").
		Where("shortURL = ?", shortURL).
		ToSql()

	var originalURL string
	row := s.Pool.QueryRow(ctx, sql, args...)
	err := row.Scan(&originalURL)
	if err != nil {
		return "", fmt.Errorf("ShortenerRepo - GetOriginalURL - row.Scan: %w", err)
	}

	return originalURL, nil
}
