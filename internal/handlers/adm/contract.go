//go:generate mockgen -destination=mock_contract_test.go -package=${GOPACKAGE} -source=contract.go
package adm

import (
	"context"

	api "github.com/s21platform/staff-service/pkg/staff/v0"
)

type StaffClient interface {
	StaffLogin(ctx context.Context, in *api.LoginRequest) (*api.LoginResponse, error)
	CreateStaff(ctx context.Context, in *api.CreateStaffRequest) (*api.Staff, error)
	ListStaff(ctx context.Context, in *api.ListStaffRequest) (*api.ListStaffResponse, error)
}
