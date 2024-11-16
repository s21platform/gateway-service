package society

import (
	"fmt"
	"net/http"

	society_proto "github.com/s21platform/society-proto/society-proto"
)

type UseCase struct {
	sC SocietyClient
}

func New(sC SocietyClient) *UseCase {
	return &UseCase{sC: sC}
}

func (u *UseCase) CreateSociety(r *http.Request) (*society_proto.SetSocietyOut, error) {
	resp, err := u.sC.CreateSociety(r.Context())
	if err != nil {
		return nil, fmt.Errorf("u.sC.CreateSociety: %v", err)
	}
	return resp, nil
}
