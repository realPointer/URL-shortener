package service

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	mock_repo "github.com/realPointer/url-shortener/internal/repo/mocks"
)

// func TestShortenerService_GetShortURL(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	type inputArgs struct {
// 		ctx         context.Context
// 		originalURL string
// 	}

// 	type getOriginalURLRepoArgs struct {
// 		ctx          context.Context
// 		shortenedURL string
// 		err          error
// 	}

// 	testCases := []struct {
// 		name             string
// 		input            inputArgs
// 		originalRepoArgs getOriginalURLRepoArgs
// 		output           string
// 		err              error
// 		wantErr          bool
// 	}{
// 		{
// 			name: "ShortURL already exists",
// 			input: inputArgs{
// 				ctx:         context.Background(),
// 				originalURL: "https://www.google.com",
// 			},
// 			originalRepoArgs: getOriginalURLRepoArgs{},
// 			output:           "123",
// 			err:              nil,
// 			wantErr:          false,
// 		},
// 		{
// 			name: "ShortURL does not exist",
// 			input: inputArgs{
// 				ctx:         context.Background(),
// 				originalURL: "https://www.google.com",
// 			},
// 			originalRepoArgs: getOriginalURLRepoArgs{
// 				ctx:          context.Background(),
// 				shortenedURL: "123",
// 				err:          errors.New("some error"),
// 			},
// 			output:  "123",
// 			err:     nil,
// 			wantErr: false,
// 		},
// 	}

// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			mockShortener := mock_repo.NewMockShortener(ctrl)
// 			mockShortener.EXPECT().GetShortURL(tc.input.ctx, tc.input.originalURL).Return(tc.output, tc.err)

// 			if tc.originalRepoArgs.err != nil {
// 				mockShortener.EXPECT().GetOriginalURL(tc.originalRepoArgs.ctx, tc.originalRepoArgs.shortenedURL).Return("", tc.originalRepoArgs.err)
// 			}

// 			shortenerService := NewShortenerService(mockShortener)

// 			shortURL, err := shortenerService.GetShortURL(tc.input.ctx, tc.input.originalURL)
// 			if tc.wantErr {
// 				assert.Error(t, err)
// 				return
// 			}

// 			assert.Equal(t, tc.output, shortURL)
// 		})
// 	}
// }

func TestShortenerService_GetOriginalURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		ctx      context.Context
		shortURL string
	}

	testCases := []struct {
		name    string
		input   args
		output  string
		err     error
		wantErr bool
	}{
		{
			name: "success",
			input: args{
				ctx:      context.Background(),
				shortURL: "123",
			},
			output:  "https://www.google.com",
			err:     nil,
			wantErr: false,
		},
		{
			name: "error",
			input: args{
				ctx:      context.Background(),
				shortURL: "123",
			},
			output:  "",
			err:     errors.New("some error"),
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockShortener := mock_repo.NewMockShortener(ctrl)
			mockShortener.EXPECT().GetOriginalURL(tc.input.ctx, tc.input.shortURL).Return(tc.output, tc.err)

			shortenerService := NewShortenerService(mockShortener)

			originalURL, err := shortenerService.GetOriginalURL(tc.input.ctx, tc.input.shortURL)
			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.Equal(t, tc.output, originalURL)
		})
	}
}
