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
		return nil, fmt.Errorf("failed to get os list in usecase: %w", err)
	}

	if resp.Options == nil {
		resp.Options = []*optionhub.Record{}
	}

	//var res model.OptionsStruct
	//res.Options = []model.Option{}
	//for _, obj := range resp.Options {
	//	res.Options = append(res.Options, model.Option{
	//		Id:    obj.Id,
	//		Label: obj.Label,
	//	})
	//}

	return resp, nil
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
