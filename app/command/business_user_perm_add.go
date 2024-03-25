package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.business/domains/business"
)

type BusinessUserPermAddCmd struct {
	Permission     string `json:"permission" validate:"required"`
	BusinessName   string
	UserName       string
	AccessUserUUID string
	AccessUserName string
}

type BusinessUserPermAddRes struct{}

type BusinessUserPermAddHandler cqrs.HandlerFunc[BusinessUserPermAddCmd, *BusinessUserPermAddRes]

func NewBusinessUserPermAddHandler(repo business.Repository, factory business.Factory, events business.Events) BusinessUserPermAddHandler {
	return func(ctx context.Context, cmd BusinessUserPermAddCmd) (*BusinessUserPermAddRes, *i18np.Error) {
		res, _err := repo.GetWithUserName(ctx, cmd.BusinessName, cmd.UserName)
		if _err != nil {
			return nil, _err
		}
		err := repo.AddUserPermission(ctx, cmd.BusinessName, business.UserDetail{
			Name: cmd.UserName,
		}, cmd.Permission)
		if err != nil {
			return nil, factory.Errors.Failed(err.Error())
		}
		events.UserPermissionAdded(&business.EventBusinessPermissionAdded{
			BusinessName:   cmd.BusinessName,
			AccessUserUUID: cmd.AccessUserUUID,
			AccessUserName: cmd.AccessUserName,
			PermissionName: cmd.Permission,
			User: business.EventUser{
				Name: cmd.UserName,
				UUID: res.User.UUID,
			},
		})
		return &BusinessUserPermAddRes{}, nil
	}
}
