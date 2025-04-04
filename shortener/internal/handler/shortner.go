package handler

import (
	"context"
	"fmt"

	shortenerpb "github.com/verestov/go-url-shortner.git/gen"
)

type Server struct {
	shortenerpb.UnimplementedShortenerServiceServer
}

// Метод для создания короткой ссылки
func (s *Server) CreateShortLink(ctx context.Context, req *shortenerpb.CreateShortLinkRequest) (*shortenerpb.ShortLinkResponse, error) {
	shortURL := fmt.Sprintf("short.ly/%s", req.GetOriginalUrl())
	return &shortenerpb.ShortLinkResponse{ShortUrl: shortURL}, nil
}

// Метод для получения оригинальной ссылки по короткой
func (s *Server) GetOriginalUrl(ctx context.Context, req *shortenerpb.GetOriginalUrlRequest) (*shortenerpb.OriginalUrlResponse, error) {
	originalURL := fmt.Sprintf("http://original.url/%s", req.GetShortUrl())
	return &shortenerpb.OriginalUrlResponse{OriginalUrl: originalURL}, nil
}
