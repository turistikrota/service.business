package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.business/domains/business"
)

type BusinessEnableCmd struct {
	BusinessName string
	UserName     string
	UserUUID     string
}

type BusinessEnableRes struct{}

type BusinessEnableHandler cqrs.HandlerFunc[BusinessEnableCmd, *BusinessEnableRes]

func NewBusinessEnableHandler(repo business.Repository, factory business.Factory, events business.Events) BusinessEnableHandler {
	return func(ctx context.Context, cmd BusinessEnableCmd) (*BusinessEnableRes, *i18np.Error) {
		err := repo.Enable(ctx, cmd.BusinessName)
		if err != nil {
			return nil, err
		}
		events.Enabled(&business.EventBusinessEnabled{
			BusinessName: cmd.BusinessName,
			UserName:     cmd.UserName,
			UserUUID:     cmd.UserUUID,
		})
		return &BusinessEnableRes{}, nil
	}
}
