package youtube

import (
	"context"
	"net/http"
)

type Service interface {
	FindStreams(ctx context.Context, videoID string) (StreamsResponse, error)
}

type service struct {
	baseURL string
	http    *http.Client
}

var _ Service = service{}

func NewService(baseURL string, httpClient *http.Client) Service {
	return service{
		baseURL,
		httpClient,
	}
}
