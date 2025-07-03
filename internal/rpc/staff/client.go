package staff

import (
	"context"
	"log"

	"github.com/s21platform/staff-service/pkg/staff"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	"github.com/s21platform/gateway-service/internal/config"
)

type Client struct {
	client staff.StaffServiceClient
}

func New(cfg *config.Config) *Client {
	conn, err := grpc.NewClient(cfg.Staff.Host+":"+cfg.Staff.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to staff service: %v", err)
	}
	client := staff.NewStaffServiceClient(conn)
	return &Client{client: client}
}

func (c *Client) StaffLogin(ctx context.Context, in *staff.LoginIn) (*staff.LoginOut, error) {
	return c.client.Login(ctx, in)
}

func (c *Client) CreateStaff(ctx context.Context, in *staff.CreateIn) (*staff.Staff, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("authorization", ctx.Value(config.KeyStaffUUID).(string)))
	resp, err := c.client.Create(ctx, in)
	if err != nil {
		log.Printf("failed to create staff: %v", err)
		return nil, err
	}
	return resp.GetStaff(), nil
}

func (c *Client) ListStaff(ctx context.Context, in *staff.ListIn) (*staff.ListOut, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("authorization", ctx.Value(config.KeyStaffUUID).(string)))
	resp, err := c.client.List(ctx, in)
	if err != nil {
		log.Printf("failed to list staff: %v", err)
		return nil, err
	}
	return resp, nil
}

func (c *Client) GetStaff(ctx context.Context, in *staff.GetIn) (*staff.Staff, error) {
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("authorization", ctx.Value(config.KeyStaffUUID).(string)))
	resp, err := c.client.Get(ctx, in)
	if err != nil {
		log.Printf("failed to get staff: %v", err)
		return nil, err
	}
	return resp.GetStaff(), nil
}
