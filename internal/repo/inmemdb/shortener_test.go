package inmemdb

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortenerRepo_SaveURLs(t *testing.T) {
	type args struct {
		ctx         context.Context
		shortURL    string
		originalURL string
	}

	testCases := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				ctx:         context.Background(),
				shortURL:    "abc",
				originalURL: "https://www.google.com",
			},
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := NewShortenerRepo()

			_, err := s.SaveURLs(tc.args.ctx, tc.args.shortURL, tc.args.originalURL)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)

			short, ok := s.originalStorage[tc.args.originalURL]
			assert.True(t, ok)
			assert.Equal(t, tc.args.shortURL, short)

			original, ok := s.shortStorage[tc.args.shortURL]
			assert.True(t, ok)
			assert.Equal(t, tc.args.originalURL, original)
		})
	}
}

func TestShortenerRepo_GetShortURL(t *testing.T) {
	type args struct {
		ctx         context.Context
		shortURL    string
		originalURL string
	}

	testCases := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				ctx:         context.Background(),
				shortURL:    "abc",
				originalURL: "https://www.google.com",
			},
			want:    "abc",
			wantErr: false,
		},
		{
			name: "Not found",
			args: args{
				ctx:         context.Background(),
				shortURL:    "abc",
				originalURL: "https://www.google.com",
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := NewShortenerRepo()

			if tc.wantErr {
				_, err := s.GetShortURL(tc.args.ctx, tc.args.originalURL)
				assert.Error(t, err)
				return
			}

			_, _ = s.SaveURLs(tc.args.ctx, tc.args.shortURL, tc.args.originalURL)

			short, err := s.GetShortURL(tc.args.ctx, tc.args.originalURL)
			if err != nil {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)

			assert.Equal(t, tc.want, short)
		})
	}
}

func TestShortenerRepo_GetOriginalURL(t *testing.T) {
	type args struct {
		ctx         context.Context
		shortURL    string
		originalURL string
	}

	testCases := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				ctx:         context.Background(),
				shortURL:    "abc",
				originalURL: "https://www.google.com",
			},
			want:    "https://www.google.com",
			wantErr: false,
		},
		{
			name: "Not found",
			args: args{
				ctx:         context.Background(),
				shortURL:    "abc",
				originalURL: "https://www.google.com",
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := NewShortenerRepo()

			if tc.wantErr {
				_, err := s.GetOriginalURL(tc.args.ctx, tc.args.shortURL)
				assert.Error(t, err)
				return
			}

			_, _ = s.SaveURLs(tc.args.ctx, tc.args.shortURL, tc.args.originalURL)

			original, err := s.GetOriginalURL(tc.args.ctx, tc.args.shortURL)
			if err != nil {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)

			assert.Equal(t, tc.want, original)
		})
	}
}
