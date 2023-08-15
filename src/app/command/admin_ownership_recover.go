package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.owner/src/domain/owner"
	"github.com/turistikrota/service.shared/decorator"
)

type AdminOwnershipRecoverCommand struct {
	OwnerNickName string
	AdminUUID     string
	Reason        string
}

type AdminOwnershipRecoverResult struct{}

type AdminOwnershipRecoverHandler decorator.CommandHandler[AdminOwnershipRecoverCommand, *AdminOwnershipRecoverResult]

type adminOwnershipRecoverHandler struct {
	repo    owner.Repository
	factory owner.Factory
	events  owner.Events
}

type AdminOwnershipRecoverConfig struct {
	Repo     owner.Repository
	Factory  owner.Factory
	Events   owner.Events
	CqrsBase decorator.Base
}

func NewAdminOwnershipRecoverHandler(config AdminOwnershipRecoverConfig) AdminOwnershipRecoverHandler {
	return decorator.ApplyCommandDecorators[AdminOwnershipRecoverCommand, *AdminOwnershipRecoverResult](
		&adminOwnershipRecoverHandler{
			repo:    config.Repo,
			factory: config.Factory,
			events:  config.Events,
		},
		config.CqrsBase,
	)
}

func (h *adminOwnershipRecoverHandler) Handle(ctx context.Context, cmd AdminOwnershipRecoverCommand) (*AdminOwnershipRecoverResult, *i18np.Error) {
	err := h.repo.Recover(ctx, cmd.OwnerNickName)
	if err != nil {
		return nil, err
	}
	h.events.RecoverByAdmin(&owner.EventOwnerRecoverByAdmin{
		OwnerNickName: cmd.OwnerNickName,
		Reason:        cmd.Reason,
		AdminUUID:     cmd.AdminUUID,
	})
	return &AdminOwnershipRecoverResult{}, nil
}
