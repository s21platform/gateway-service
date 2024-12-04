package option

import (
	"fmt"
	"net/http"

	"github.com/s21platform/gateway-service/internal/model"

	optionhub "github.com/s21platform/optionhub-proto/optionhub-proto"
)

type Usecase struct {
	oC OptionClient
}

func New(oC OptionClient) *Usecase {
	return &Usecase{oC: oC}
}

func (uc *Usecase) GetOsList(r *http.Request) (*model.OptionsStruct, error) {
	name := r.URL.Query().Get("name")
	searchName := &optionhub.GetByNameIn{Name: name}

	resp, err := uc.oC.GetOsBySearchName(r.Context(), searchName)
	if err != nil {
		return nil, fmt.Errorf("failed to get os list in usecase: %w", err)
	}

	var res model.OptionsStruct
	for _, obj := range resp.Options {
		res.Options = append(res.Options, model.Option{
			Id:    obj.Id,
			Label: obj.Label,
		})
	}

	return &res, nil
}

func (uc *Usecase) GetWorkPlaceList(r *http.Request) (*optionhub.GetByNameOut, error) {
	name := r.URL.Query().Get("name")
	searchName := &optionhub.GetByNameIn{Name: name}

	resp, err := uc.oC.GetWorkPlaceBySearchName(r.Context(), searchName)
	if err != nil {
		return nil, fmt.Errorf("failed to get workplace list in usecase: %w", err)
	}

	return resp, nil
}

func (uc *Usecase) GetStudyPlaceList(r *http.Request) (*optionhub.GetByNameOut, error) {
	name := r.URL.Query().Get("name")
	searchName := &optionhub.GetByNameIn{Name: name}

	resp, err := uc.oC.GetStudyPlaceBySearchName(r.Context(), searchName)
	if err != nil {
		return nil, fmt.Errorf("failed to get study place list in usecase: %w", err)
	}

	return resp, nil
}

func (uc *Usecase) GetHobbyList(r *http.Request) (*optionhub.GetByNameOut, error) {
	name := r.URL.Query().Get("name")
	searchName := &optionhub.GetByNameIn{Name: name}

	resp, err := uc.oC.GetHobbyBySearchName(r.Context(), searchName)
	if err != nil {
		return nil, fmt.Errorf("failed to get hobby list in usecase: %w", err)
	}

	return resp, nil
}

func (uc *Usecase) GetSkillList(r *http.Request) (*optionhub.GetByNameOut, error) {
	name := r.URL.Query().Get("name")
	searchName := &optionhub.GetByNameIn{Name: name}

	resp, err := uc.oC.GetSkillBySearchName(r.Context(), searchName)
	if err != nil {
		return nil, fmt.Errorf("failed to get skill list in usecase: %w", err)
	}

	return resp, nil
}

func (uc *Usecase) GetCityList(r *http.Request) (*optionhub.GetByNameOut, error) {
	name := r.URL.Query().Get("name")
	searchName := &optionhub.GetByNameIn{Name: name}

	resp, err := uc.oC.GetCityBySearchName(r.Context(), searchName)
	if err != nil {
		return nil, fmt.Errorf("failed to get city list in usecase: %w", err)
	}

	return resp, nil
}

func (uc *Usecase) GetSocietyDirectionList(r *http.Request) (*optionhub.GetByNameOut, error) {
	name := r.URL.Query().Get("name")
	searchName := &optionhub.GetByNameIn{Name: name}

	resp, err := uc.oC.GetSocietyDirectionBySearchName(r.Context(), searchName)
	if err != nil {
		return nil, fmt.Errorf("failed to get society direction list in usecase: %w", err)
	}

	return resp, nil
}
