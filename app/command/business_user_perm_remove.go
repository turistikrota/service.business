package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.business/domains/business"
)

type BusinessUserPermRemoveCmd struct {
	Permission     string `json:"permission" validate:"required"`
	BusinessName   string `params:"-" json:"-"`
	UserName       string `params:"userName" validate:"required,username"`
	AccessUserUUID string `params:"-" json:"-"`
	AccessUserName string `params:"-" json:"-"`
}

type BusinessUserPermRemoveRes struct{}

type BusinessUserPermRemoveHandler cqrs.HandlerFunc[BusinessUserPermRemoveCmd, *BusinessUserPermRemoveRes]

func NewBusinessUserPermRemoveHandler(repo business.Repository, factory business.Factory, events business.Events) BusinessUserPermRemoveHandler {
	return func(ctx context.Context, cmd BusinessUserPermRemoveCmd) (*BusinessUserPermRemoveRes, *i18np.Error) {
		res, _err := repo.GetWithUserName(ctx, cmd.BusinessName, cmd.UserName)
		if _err != nil {
			return nil, _err
		}
		err := repo.RemoveUserPermission(ctx, cmd.BusinessName, business.UserDetail{
			Name: cmd.UserName,
		}, cmd.Permission)
		if err != nil {
			return nil, factory.Errors.Failed(err.Error())
		}
		events.UserPermissionRemoved(&business.EventBusinessPermissionRemoved{
			BusinessName:   cmd.BusinessName,
			AccessUserUUID: cmd.AccessUserUUID,
			AccessUserName: cmd.AccessUserName,
			PermissionName: cmd.Permission,
			User: business.EventUser{
				Name: cmd.UserName,
				UUID: res.User.UUID,
			},
		})
		return &BusinessUserPermRemoveRes{}, nil
	}
}
