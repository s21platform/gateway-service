package option

import (
	"fmt"
	"net/http"

	optionhub "github.com/s21platform/optionhub-proto/optionhub-proto"
)

type Usecase struct {
	oC OptionClient
}

func New(oC OptionClient) *Usecase {
	return &Usecase{oC: oC}
}

func (uc *Usecase) GetOsList(r *http.Request) (*optionhub.GetByNameOut, error) {
	name := r.URL.Query().Get("name")
	searchName := &optionhub.GetByNameIn{Name: name}

	resp, err := uc.oC.GetOsBySearchName(r.Context(), searchName)
	if err != nil {
		return nil, fmt.Errorf("failed to get os list in usercase: %w", err)
	}

	return resp, nil
}

func (uc *Usecase) GetSocietyDirectionList(r *http.Request) (*optionhub.GetByNameOut, error) {
	name := r.URL.Query().Get("name")
	searchName := &optionhub.GetByNameIn{Name: name}

	resp, err := uc.oC.GetSocietyDirectionBySearchName(r.Context(), searchName)
	if err != nil {
		return nil, fmt.Errorf("failed to get society direction list in usercase: %w", err)
	}

	return resp, nil
}
