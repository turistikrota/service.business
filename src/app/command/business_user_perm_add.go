package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.business/src/domain/business"
)

type BusinessUserPermAddCommand struct {
	BusinessNickName string
	UserName         string
	AccessUserUUID   string
	AccessUserName   string
	PermissionName   string
}

type BusinessUserPermAddResult struct{}

type BusinessUserPermAddHandler decorator.CommandHandler[BusinessUserPermAddCommand, *BusinessUserPermAddResult]

type businessUserPermAddHandler struct {
	repo    business.Repository
	factory business.Factory
	events  business.Events
}

type BusinessUserPermAddHandlerConfig struct {
	Repo     business.Repository
	Factory  business.Factory
	Events   business.Events
	CqrsBase decorator.Base
}

func NewBusinessUserPermAddHandler(config BusinessUserPermAddHandlerConfig) BusinessUserPermAddHandler {
	return decorator.ApplyCommandDecorators[BusinessUserPermAddCommand, *BusinessUserPermAddResult](
		&businessUserPermAddHandler{
			repo:    config.Repo,
			factory: config.Factory,
			events:  config.Events,
		},
		config.CqrsBase,
	)
}

func (h *businessUserPermAddHandler) Handle(ctx context.Context, cmd BusinessUserPermAddCommand) (*BusinessUserPermAddResult, *i18np.Error) {
	res, _err := h.repo.GetWithUserName(ctx, cmd.BusinessNickName, cmd.UserName)
	if _err != nil {
		return nil, _err
	}
	err := h.repo.AddUserPermission(ctx, cmd.BusinessNickName, business.UserDetail{
		Name: cmd.UserName,
	}, cmd.PermissionName)
	if err != nil {
		return nil, h.factory.Errors.Failed(err.Error())
	}
	h.events.UserPermissionAdded(&business.EventBusinessPermissionAdded{
		BusinessUUID:   res.Entity.UUID,
		AccessUserUUID: cmd.AccessUserUUID,
		AccessUserName: cmd.AccessUserName,
		PermissionName: cmd.PermissionName,
		User: business.EventUser{
			Name: cmd.UserName,
			UUID: res.User.UUID,
		},
	})
	return &BusinessUserPermAddResult{}, nil
}
