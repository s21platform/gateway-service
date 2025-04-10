//go:generate mockgen -destination=mock_contract_test.go -package=${GOPACKAGE} -source=contract.go
package adm

import (
	"net/http"

	"github.com/s21platform/staff-service/pkg/staff"
)

type StaffService interface {
	StaffLogin(r *http.Request) (*staff.LoginOut, error)
	CreateStaff(r *http.Request) (*staff.Staff, error)
	ListStaff(r *http.Request) (*staff.ListOut, error)
	GetStaff(r *http.Request, staffID string) (*staff.Staff, error)
}
