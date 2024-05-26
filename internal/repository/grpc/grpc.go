package grpc

import (
	"context"
	auth_proto "github.com/s21platform/auth-proto/auth-proto"
	"google.golang.org/grpc"
)

type ServiceClients struct {
	AuthClient auth_proto.AuthServiceClient
}

func (g *ServiceClients) Login(ctx context.Context, in *auth_proto.LoginRequest, opts ...grpc.CallOption) (*auth_proto.LoginResponse, error) {
	resp, err := g.AuthClient.Login(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
