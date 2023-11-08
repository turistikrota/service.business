package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.owner/src/domain/owner"
)

type AdminOwnershipRecoverCommand struct {
	OwnerNickName string
	AdminUUID     string
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
	res, _err := h.repo.GetByNickName(ctx, cmd.OwnerNickName)
	if _err != nil {
		return nil, _err
	}
	err := h.repo.Recover(ctx, cmd.OwnerNickName)
	if err != nil {
		return nil, err
	}
	h.events.RecoverByAdmin(&owner.EventOwnerRecoverByAdmin{
		OwnerUUID: res.UUID,
		AdminUUID: cmd.AdminUUID,
	})
	return &AdminOwnershipRecoverResult{}, nil
}
