package search

import (
	"context"
	"fmt"
	"github.com/s21platform/gateway-service/internal/config"
	"github.com/s21platform/search-proto/search"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type Client struct {
	client search.SearchServiceClient
}

func New(cfg *config.Config) *Client {
	connStr := fmt.Sprintf("%s:%s", cfg.Search.Host, cfg.Search.Port)
	conn, err := grpc.NewClient(connStr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("faild to connect to search service: %v", err)
	}
	client := search.NewSearchServiceClient(conn)
	return &Client{client: client}
}

func (c *Client) SearchSociety(ctx context.Context, searchName string) (*search.GetSocietyOut, error) {
	result, err := c.client.GetSociety(ctx, &search.GetSocietyIn{
		PartName: searchName,
	})
	if err != nil {
		return nil, fmt.Errorf("faild to search society: %v", err)
	}
	return result, nil
}
