package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.owner/src/domain/owner"
)

type OwnershipUserRemoveCommand struct {
	OwnerNickName  string
	UserName       string
	AccessUserUUID string
}

type OwnershipUserRemoveResult struct{}

type OwnershipUserRemoveHandler decorator.CommandHandler[OwnershipUserRemoveCommand, *OwnershipUserRemoveResult]

type ownershipUserRemoveHandler struct {
	repo    owner.Repository
	factory owner.Factory
	events  owner.Events
}

type OwnershipUserRemoveHandlerConfig struct {
	Repo     owner.Repository
	Factory  owner.Factory
	Events   owner.Events
	CqrsBase decorator.Base
}

func NewOwnershipUserRemoveHandler(config OwnershipUserRemoveHandlerConfig) OwnershipUserRemoveHandler {
	return decorator.ApplyCommandDecorators[OwnershipUserRemoveCommand, *OwnershipUserRemoveResult](
		&ownershipUserRemoveHandler{
			repo:    config.Repo,
			factory: config.Factory,
			events:  config.Events,
		},
		config.CqrsBase,
	)
}

func (h *ownershipUserRemoveHandler) Handle(ctx context.Context, cmd OwnershipUserRemoveCommand) (*OwnershipUserRemoveResult, *i18np.Error) {
	res, _err := h.repo.GetByNickName(ctx, cmd.OwnerNickName)
	if _err != nil {
		return nil, _err
	}
	err := h.repo.RemoveUser(ctx, cmd.OwnerNickName, owner.UserDetail{
		Name: cmd.UserName,
	})
	if err != nil {
		return nil, h.factory.Errors.Failed("failed to remove user from ownership")
	}
	h.events.UserRemoved(&owner.EventOwnerUserRemoved{
		OwnerUUID:      res.UUID,
		AccessUserUUID: cmd.AccessUserUUID,
		User: owner.EventUser{
			Name: cmd.UserName,
		},
	})
	return &OwnershipUserRemoveResult{}, nil
}
