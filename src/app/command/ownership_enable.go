package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.owner/src/domain/owner"
	"github.com/turistikrota/service.shared/decorator"
)

type OwnershipEnableCommand struct {
	OwnerNickName string
	UserName      string
	UserCode      string
	UserUUID      string
}

type OwnershipEnableResult struct{}

type OwnershipEnableHandler decorator.CommandHandler[OwnershipEnableCommand, *OwnershipEnableResult]

type ownershipEnableHandler struct {
	repo    owner.Repository
	factory owner.Factory
	events  owner.Events
}

type OwnershipEnableConfig struct {
	Repo     owner.Repository
	Factory  owner.Factory
	Events   owner.Events
	CqrsBase decorator.Base
}

func NewOwnershipEnableHandler(config OwnershipEnableConfig) OwnershipEnableHandler {
	return decorator.ApplyCommandDecorators[OwnershipEnableCommand, *OwnershipEnableResult](
		&ownershipEnableHandler{
			repo:    config.Repo,
			factory: config.Factory,
			events:  config.Events,
		},
		config.CqrsBase,
	)
}

func (h *ownershipEnableHandler) Handle(ctx context.Context, cmd OwnershipEnableCommand) (*OwnershipEnableResult, *i18np.Error) {
	err := h.repo.Enable(ctx, cmd.OwnerNickName)
	if err != nil {
		return nil, err
	}
	h.events.Enabled(&owner.EventOwnerEnabled{
		OwnerNickName: cmd.OwnerNickName,
		UserName:      cmd.UserName,
		UserCode:      cmd.UserCode,
		UserUUID:      cmd.UserUUID,
	})
	return &OwnershipEnableResult{}, nil
}
