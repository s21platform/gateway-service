//go:generate mockgen -source $GOFILE -destination mock_contract_test.go -package $GOPACKAGE
package RESTHandlers

import (
	"context"
	auth_proto "github.com/s21platform/auth-proto/auth-proto"
	"google.golang.org/grpc"
)

type Auth interface {
	Login(ctx context.Context, in *auth_proto.LoginRequest, opts ...grpc.CallOption) (*auth_proto.LoginResponse, error)
}
