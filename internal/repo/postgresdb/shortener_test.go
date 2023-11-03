package postgresrepo

import (
	"context"
	"errors"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/pashagolub/pgxmock/v3"
	"github.com/realPointer/url-shortener/pkg/postgres"
	"github.com/stretchr/testify/assert"
)

func TestShortenerRepo_SaveURLs(t *testing.T) {
	type args struct {
		ctx         context.Context
		shortURL    string
		originalURL string
	}

	type MockBehavior func(m pgxmock.PgxPoolIface, args args)

	testCases := []struct {
		name         string
		args         args
		mockBehavior MockBehavior
		want         string
		wantErr      bool
	}{
		{
			name: "OK",
			args: args{
				ctx:         context.Background(),
				shortURL:    "abc",
				originalURL: "https://www.google.com",
			},
			mockBehavior: func(m pgxmock.PgxPoolIface, args args) {
				m.ExpectQuery("INSERT INTO links").
					WithArgs(args.shortURL, args.originalURL).
					WillReturnRows(pgxmock.NewRows([]string{"shortURL"}).AddRow(args.shortURL))
			},
			want:    "abc",
			wantErr: false,
		},
		{
			name: "Error",
			args: args{
				ctx:         context.Background(),
				shortURL:    "abc",
				originalURL: "https://www.google.com",
			},
			mockBehavior: func(m pgxmock.PgxPoolIface, args args) {
				m.ExpectQuery("INSERT INTO links").
					WithArgs(args.shortURL, args.originalURL).
					WillReturnError(errors.New("something went wrong"))
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			poolMock, _ := pgxmock.NewPool()
			defer poolMock.Close()
			tc.mockBehavior(poolMock, tc.args)

			postgresMock := &postgres.Postgres{
				Builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
				Pool:    poolMock,
			}
			shortenerRepoMock := NewShortenerRepo(postgresMock)

			shortURL, err := shortenerRepoMock.SaveURLs(tc.args.ctx, tc.args.shortURL, tc.args.originalURL)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.want, shortURL)

			err = poolMock.ExpectationsWereMet()
			assert.NoError(t, err)
		})
	}
}

func TestShortenerRepo_GetShortURL(t *testing.T) {
	type args struct {
		ctx         context.Context
		originalURL string
	}

	type MockBehavior func(m pgxmock.PgxPoolIface, args args)

	testCases := []struct {
		name         string
		args         args
		mockBehavior MockBehavior
		want         string
		wantErr      bool
	}{
		{
			name: "OK",
			args: args{
				ctx:         context.Background(),
				originalURL: "https://www.google.com",
			},
			mockBehavior: func(m pgxmock.PgxPoolIface, args args) {
				m.ExpectQuery("SELECT shortURL FROM links").
					WithArgs(args.originalURL).
					WillReturnRows(pgxmock.NewRows([]string{"shortURL"}).AddRow("abc"))
			},
			want:    "abc",
			wantErr: false,
		},
		{
			name: "Error",
			args: args{
				ctx:         context.Background(),
				originalURL: "https://www.google.com",
			},
			mockBehavior: func(m pgxmock.PgxPoolIface, args args) {
				m.ExpectQuery("SELECT shortURL FROM links").
					WithArgs(args.originalURL).
					WillReturnError(errors.New("something went wrong"))
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			poolMock, _ := pgxmock.NewPool()
			defer poolMock.Close()
			tc.mockBehavior(poolMock, tc.args)

			postgresMock := &postgres.Postgres{
				Builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
				Pool:    poolMock,
			}

			shortenerRepoMock := NewShortenerRepo(postgresMock)

			shortURL, err := shortenerRepoMock.GetShortURL(tc.args.ctx, tc.args.originalURL)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.want, shortURL)

			err = poolMock.ExpectationsWereMet()
			assert.NoError(t, err)
		})
	}
}

func TestShortenerRepo_GetOriginalURL(t *testing.T) {
	type args struct {
		ctx      context.Context
		shortURL string
	}

	type MockBehavior func(m pgxmock.PgxPoolIface, args args)

	testCases := []struct {
		name         string
		args         args
		mockBehavior MockBehavior
		want         string
		wantErr      bool
	}{
		{
			name: "OK",
			args: args{
				ctx:      context.Background(),
				shortURL: "abc",
			},
			mockBehavior: func(m pgxmock.PgxPoolIface, args args) {
				m.ExpectQuery("SELECT originalURL FROM links").
					WithArgs(args.shortURL).
					WillReturnRows(pgxmock.NewRows([]string{"originalURL"}).AddRow("https://www.google.com"))
			},
			want:    "https://www.google.com",
			wantErr: false,
		},
		{
			name: "Error",
			args: args{
				ctx:      context.Background(),
				shortURL: "abc",
			},
			mockBehavior: func(m pgxmock.PgxPoolIface, args args) {
				m.ExpectQuery("SELECT originalURL FROM links").
					WithArgs(args.shortURL).
					WillReturnError(errors.New("something went wrong"))
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			poolMock, _ := pgxmock.NewPool()
			defer poolMock.Close()
			tc.mockBehavior(poolMock, tc.args)

			postgresMock := &postgres.Postgres{
				Builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
				Pool:    poolMock,
			}

			shortenerRepoMock := NewShortenerRepo(postgresMock)

			originalURL, err := shortenerRepoMock.GetOriginalURL(tc.args.ctx, tc.args.shortURL)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.want, originalURL)

			err = poolMock.ExpectationsWereMet()
			assert.NoError(t, err)
		})
	}
}
