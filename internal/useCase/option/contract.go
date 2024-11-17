package option

import (
	"context"

	optionhub "github.com/s21platform/optionhub-proto/optionhub-proto"
)

type OptionClient interface {
	GetOSById(ctx context.Context, id int64) (*optionhub.GetByIdOut, error)
}
