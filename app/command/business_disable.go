package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.business/domains/business"
)

type BusinessDisableCmd struct {
	BusinessName string
	UserName     string
	UserUUID     string
}

type BusinessDisableRes struct{}

type BusinessDisableHandler cqrs.HandlerFunc[BusinessDisableCmd, *BusinessDisableRes]

func NewBusinessDisableHandler(repo business.Repository, factory business.Factory, events business.Events) BusinessDisableHandler {
	return func(ctx context.Context, cmd BusinessDisableCmd) (*BusinessDisableRes, *i18np.Error) {
		err := repo.Disable(ctx, cmd.BusinessName)
		if err != nil {
			return nil, err
		}
		events.Disabled(&business.EventBusinessDisabled{
			BusinessName: cmd.BusinessName,
			UserName:     cmd.UserName,
			UserUUID:     cmd.UserUUID,
		})
		return &BusinessDisableRes{}, nil
	}
}
