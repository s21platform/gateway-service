package staff

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	"github.com/s21platform/gateway-service/internal/config"
	api "github.com/s21platform/staff-service/pkg/staff/v0"
)

type Client struct {
	client api.StaffServiceClient
}

func New(cfg *config.Config) *Client {
	conn, err := grpc.NewClient(cfg.Staff.Host+":"+cfg.Staff.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to staff service: %v", err)
	}
	client := api.NewStaffServiceClient(conn)
	return &Client{client: client}
}

func (c *Client) StaffLogin(ctx context.Context, in *api.LoginRequest) (*api.LoginResponse, error) {
	return c.client.Login(ctx, in)
}

func (c *Client) CreateStaff(ctx context.Context, in *api.CreateStaffRequest) (*api.Staff, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("authorization", ctx.Value(config.KeyStaffUUID).(string)))
	resp, err := c.client.CreateStaff(ctx, in)
	if err != nil {
		log.Printf("failed to create staff: %v", err)
		return nil, err
	}
	return resp.GetStaff(), nil
}
