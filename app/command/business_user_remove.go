package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.business/domains/business"
)

type BusinessUserRemoveCmd struct {
	BusinessName   string `params:"-"`
	UserName       string `params:"userName" validate:"required,username"`
	AccessUserUUID string `params:"-"`
	AccessUserName string `params:"-"`
}

type BusinessUserRemoveRes struct{}

type BusinessUserRemoveHandler cqrs.HandlerFunc[BusinessUserRemoveCmd, *BusinessUserRemoveRes]

func NewBusinessUserRemoveHandler(repo business.Repository, factory business.Factory, events business.Events) BusinessUserRemoveHandler {
	return func(ctx context.Context, cmd BusinessUserRemoveCmd) (*BusinessUserRemoveRes, *i18np.Error) {
		res, _err := repo.GetWithUserName(ctx, cmd.BusinessName, cmd.UserName)
		if _err != nil {
			return nil, _err
		}
		err := repo.RemoveUser(ctx, cmd.BusinessName, business.UserDetail{
			Name: cmd.UserName,
		})
		if err != nil {
			return nil, factory.Errors.Failed(err.Error())
		}
		events.UserRemoved(&business.EventBusinessUserRemoved{
			BusinessName:   cmd.BusinessName,
			AccessUserUUID: cmd.AccessUserUUID,
			AccessUserName: cmd.AccessUserName,
			User: business.EventUser{
				Name: cmd.UserName,
				UUID: res.User.UUID,
			},
		})
		return &BusinessUserRemoveRes{}, nil
	}
}
