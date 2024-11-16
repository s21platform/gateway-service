package society

import (
	"context"

	society "github.com/s21platform/society-proto/society-proto"
)

type SocietyClient interface {
	CreateSociety(ctx context.Context, name string, desc string, isPrivate bool, dirID int64, accessLevel int64) (*society.SetSocietyOut, error)
}
