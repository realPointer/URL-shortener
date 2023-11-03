package grpccontroller

import (
	"context"

	"github.com/realPointer/url-shortener/internal/service"
	pb "github.com/realPointer/url-shortener/pkg/pb"
)

type Service struct {
	pb.UnimplementedShortenerServer
	shortenerService service.Shortener
}

func NewService(shortenerService service.Shortener) pb.ShortenerServer {
	return &Service{
		shortenerService: shortenerService,
	}
}

func (s *Service) GetOriginalURL(ctx context.Context, in *pb.GetOriginalURLRequest) (*pb.GetOriginalURLResponse, error) {
	originalURL, err := s.shortenerService.GetOriginalURL(ctx, in.ShortURL)
	if err != nil {
		// Handle error
		return nil, err
	}
	return &pb.GetOriginalURLResponse{OriginalURL: originalURL}, nil
}

func (s *Service) CreateShortURL(ctx context.Context, in *pb.CreateShortURLRequest) (*pb.CreateShortURLResponse, error) {
	shortURL, err := s.shortenerService.GetShortURL(ctx, in.OriginalURL)
	if err != nil {
		// Handle error
		return nil, err
	}
	return &pb.CreateShortURLResponse{ShortURL: shortURL}, nil
}
