package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.owner/src/domain/owner"
	"github.com/turistikrota/service.shared/decorator"
)

type OwnershipUserPermRemoveCommand struct {
	OwnerNickName  string
	UserName       string
	UserCode       string
	AccessUserUUID string
	PermissionName string
}

type OwnershipUserPermRemoveResult struct{}

type OwnershipUserPermRemoveHandler decorator.CommandHandler[OwnershipUserPermRemoveCommand, *OwnershipUserPermRemoveResult]

type ownershipUserPermRemoveHandler struct {
	repo    owner.Repository
	factory owner.Factory
	events  owner.Events
}

type OwnershipUserPermRemoveHandlerConfig struct {
	Repo     owner.Repository
	Factory  owner.Factory
	Events   owner.Events
	CqrsBase decorator.Base
}

func NewOwnershipUserPermRemoveHandler(config OwnershipUserPermRemoveHandlerConfig) OwnershipUserPermRemoveHandler {
	return decorator.ApplyCommandDecorators[OwnershipUserPermRemoveCommand, *OwnershipUserPermRemoveResult](
		&ownershipUserPermRemoveHandler{
			repo:    config.Repo,
			factory: config.Factory,
			events:  config.Events,
		},
		config.CqrsBase,
	)
}

func (h *ownershipUserPermRemoveHandler) Handle(ctx context.Context, cmd OwnershipUserPermRemoveCommand) (*OwnershipUserPermRemoveResult, *i18np.Error) {
	err := h.repo.RemoveUserPermission(ctx, cmd.OwnerNickName, owner.UserDetail{
		Name: cmd.UserName,
		Code: cmd.UserCode,
	}, cmd.PermissionName)
	if err != nil {
		return nil, h.factory.Errors.Failed("failed to remove user permission from ownership")
	}
	h.events.UserPermissionRemoved(&owner.EventOwnerPermissionRemoved{
		OwnerNickName:  cmd.OwnerNickName,
		AccessUserUUID: cmd.AccessUserUUID,
		PermissionName: cmd.PermissionName,
		User: owner.EventUser{
			Name: cmd.UserName,
			Code: cmd.UserCode,
		},
	})
	return &OwnershipUserPermRemoveResult{}, nil
}
