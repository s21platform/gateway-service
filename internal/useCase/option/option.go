package option

import (
	"fmt"
	"net/http"

	"github.com/s21platform/gateway-service/internal/config"
	optionhub "github.com/s21platform/optionhub-proto/optionhub-proto"
)

type Usecase struct {
	oC OptionClient
}

func New(oC OptionClient) *Usecase {
	return &Usecase{oC: oC}
}

func (uc *Usecase) GetOS(r *http.Request) (*optionhub.GetByIdOut, error) {
	id := r.Context().Value(config.KeyID).(int64)

	resp, err := uc.oC.GetOSByID(r.Context(), id)
	if err != nil {
		return nil, fmt.Errorf("failed to get os in usercase: %w", err)
	}

	return resp, nil
}
