package option

import (
	"context"

	optionhub "github.com/s21platform/optionhub-proto/optionhub-proto"
)

type OptionClient interface {
	GetAllOs(ctx context.Context) (*optionhub.GetAllOut, error)
}
