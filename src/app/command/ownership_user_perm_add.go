package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.owner/src/domain/owner"
	"github.com/turistikrota/service.shared/decorator"
)

type OwnershipUserPermAddCommand struct {
	OwnerNickName  string
	UserName       string
	UserCode       string
	AccessUserUUID string
	PermissionName string
}

type OwnershipUserPermAddResult struct{}

type OwnershipUserPermAddHandler decorator.CommandHandler[OwnershipUserPermAddCommand, *OwnershipUserPermAddResult]

type ownershipUserPermAddHandler struct {
	repo    owner.Repository
	factory owner.Factory
	events  owner.Events
}

type OwnershipUserPermAddHandlerConfig struct {
	Repo     owner.Repository
	Factory  owner.Factory
	Events   owner.Events
	CqrsBase decorator.Base
}

func NewOwnershipUserPermAddHandler(config OwnershipUserPermAddHandlerConfig) OwnershipUserPermAddHandler {
	return decorator.ApplyCommandDecorators[OwnershipUserPermAddCommand, *OwnershipUserPermAddResult](
		&ownershipUserPermAddHandler{
			repo:    config.Repo,
			factory: config.Factory,
			events:  config.Events,
		},
		config.CqrsBase,
	)
}

func (h *ownershipUserPermAddHandler) Handle(ctx context.Context, cmd OwnershipUserPermAddCommand) (*OwnershipUserPermAddResult, *i18np.Error) {
	err := h.repo.AddUserPermission(ctx, cmd.OwnerNickName, owner.UserDetail{
		Name: cmd.UserName,
		Code: cmd.UserCode,
	}, cmd.PermissionName)
	if err != nil {
		return nil, h.factory.Errors.Failed(err.Error())
	}
	h.events.UserPermissionAdded(&owner.EventOwnerPermissionAdded{
		OwnerNickName:  cmd.OwnerNickName,
		AccessUserUUID: cmd.AccessUserUUID,
		PermissionName: cmd.PermissionName,
		User: owner.EventUser{
			Name: cmd.UserName,
			Code: cmd.UserCode,
		},
	})
	return &OwnershipUserPermAddResult{}, nil
}
