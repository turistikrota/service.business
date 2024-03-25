package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.business/domains/business"
)

type BusinessSetLocaleCmd struct {
	BusinessName string
	Locale       string
}

type BusinessSetLocaleRes struct{}

type BusinessSetLocaleHandler cqrs.HandlerFunc[BusinessSetLocaleCmd, *BusinessSetLocaleRes]

func NewBusinessSetLocaleHandler(repo business.Repository, factory business.Factory) BusinessSetLocaleHandler {
	return func(ctx context.Context, cmd BusinessSetLocaleCmd) (*BusinessSetLocaleRes, *i18np.Error) {
		err := repo.SetPreferredLocale(ctx, cmd.BusinessName, cmd.Locale)
		if err != nil {
			return nil, err
		}
		return &BusinessSetLocaleRes{}, nil
	}
}
