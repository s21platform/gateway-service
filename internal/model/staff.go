package model

import (
	"fmt"

	"github.com/s21platform/staff-service/pkg/staff"
)

type StaffAuth struct {
	ID           string       `json:"id"`
	Login        string       `json:"login"`
	RoleID       int32        `json:"role_id"`
	RoleName     string       `json:"role_name"`
	Permissions  *Permissions `json:"permissions,omitempty"`
	CreatedAt    int64        `json:"created_at"`
	UpdatedAt    int64        `json:"updated_at"`
	AccessToken  string       `json:"access_token"`
	RefreshToken string       `json:"refresh_token"`
	ExpiresAt    int64        `json:"expires_at"`
}

type Permissions struct {
	Access []string `json:"access"`
}

// FromDTO преобразует данные из staff.LoginResponse в модель StaffAuth
func (s *StaffAuth) FromDTO(response *staff.LoginOut) error {
	if response == nil {
		return fmt.Errorf("response is nil")
	}

	staff := response.GetStaff()
	if staff == nil {
		return fmt.Errorf("staff info is nil")
	}

	s.ID = staff.GetId()
	s.Login = staff.GetLogin()
	s.RoleID = staff.GetRoleId()
	s.RoleName = staff.GetRoleName()
	s.CreatedAt = staff.GetCreatedAt()
	s.UpdatedAt = staff.GetUpdatedAt()
	s.AccessToken = response.GetAccessToken()
	s.RefreshToken = response.GetRefreshToken()
	s.ExpiresAt = response.GetExpiresAt()

	if perms := staff.GetPermissions(); perms != nil && len(perms.GetAccess()) > 0 {
		s.Permissions = &Permissions{
			Access: perms.GetAccess(),
		}
	} else {
		s.Permissions = nil
	}

	return nil
}
