package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.owner/src/domain/owner"
)

type OwnershipDisableCommand struct {
	OwnerNickName string
	UserName      string
	UserCode      string
	UserUUID      string
}

type OwnershipDisableResult struct{}

type OwnershipDisableHandler decorator.CommandHandler[OwnershipDisableCommand, *OwnershipDisableResult]

type ownershipDisableHandler struct {
	repo    owner.Repository
	factory owner.Factory
	events  owner.Events
}

type OwnershipDisableConfig struct {
	Repo     owner.Repository
	Factory  owner.Factory
	Events   owner.Events
	CqrsBase decorator.Base
}

func NewOwnershipDisableHandler(config OwnershipDisableConfig) OwnershipDisableHandler {
	return decorator.ApplyCommandDecorators[OwnershipDisableCommand, *OwnershipDisableResult](
		&ownershipDisableHandler{
			repo:    config.Repo,
			factory: config.Factory,
			events:  config.Events,
		},
		config.CqrsBase,
	)
}

func (h *ownershipDisableHandler) Handle(ctx context.Context, cmd OwnershipDisableCommand) (*OwnershipDisableResult, *i18np.Error) {
	err := h.repo.Disable(ctx, cmd.OwnerNickName)
	if err != nil {
		return nil, err
	}

	h.events.Disabled(&owner.EventOwnerDisabled{
		OwnerNickName: cmd.OwnerNickName,
		UserName:      cmd.UserName,
		UserCode:      cmd.UserCode,
		UserUUID:      cmd.UserUUID,
	})
	return &OwnershipDisableResult{}, nil
}
