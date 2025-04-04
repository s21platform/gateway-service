//go:generate mockgen -destination=mock_contract_test.go -package=${GOPACKAGE} -source=contract.go
package adm

import (
	"context"

	"github.com/s21platform/staff-service/pkg/staff"
)

type StaffClient interface {
	StaffLogin(ctx context.Context, in *staff.LoginIn) (*staff.LoginOut, error)
	CreateStaff(ctx context.Context, in *staff.CreateIn) (*staff.Staff, error)
	ListStaff(ctx context.Context, in *staff.ListIn) (*staff.ListOut, error)
}
