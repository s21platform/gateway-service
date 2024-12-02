package option

import (
	"context"

	optionhub "github.com/s21platform/optionhub-proto/optionhub-proto"
)

type OptionClient interface {
	GetOsBySearchName(ctx context.Context, searchName *optionhub.GetByNameIn) (*optionhub.GetByNameOut, error)
	GetSocietyDirectionBySearchName(ctx context.Context, searchName *optionhub.GetByNameIn) (*optionhub.GetByNameOut, error)
}
