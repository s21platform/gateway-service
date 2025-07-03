package option

import (
	"context"

	"github.com/s21platform/gateway-service/internal/model"

	optionhub "github.com/s21platform/optionhub-proto/optionhub-proto"
)

type OptionClient interface {
	GetOsBySearchName(ctx context.Context, searchName *optionhub.GetByNameIn) (*optionhub.GetByNameOut, error)
	GetWorkPlaceBySearchName(ctx context.Context, searchName *optionhub.GetByNameIn) (*optionhub.GetByNameOut, error)
	GetStudyPlaceBySearchName(ctx context.Context, searchName *optionhub.GetByNameIn) (*optionhub.GetByNameOut, error)
	GetHobbyBySearchName(ctx context.Context, searchName *optionhub.GetByNameIn) (*optionhub.GetByNameOut, error)
	GetSkillBySearchName(ctx context.Context, searchName *optionhub.GetByNameIn) (*optionhub.GetByNameOut, error)
	GetCityBySearchName(ctx context.Context, searchName *optionhub.GetByNameIn) (*optionhub.GetByNameOut, error)
	GetSocietyDirectionBySearchName(ctx context.Context, searchName *optionhub.GetByNameIn) (*optionhub.GetByNameOut, error)
	GetOptionRequests(ctx context.Context) (model.OptionRequestsList, error)
}
