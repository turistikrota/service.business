package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.owner/src/domain/owner"
)

type AdminOwnershipRejectCommand struct {
	OwnerNickName string
	AdminUUID     string
	Reason        string
}

type AdminOwnershipRejectResult struct{}

type AdminOwnershipRejectHandler decorator.CommandHandler[AdminOwnershipRejectCommand, *AdminOwnershipRejectResult]

type adminOwnershipRejectHandler struct {
	repo    owner.Repository
	factory owner.Factory
	events  owner.Events
}

type AdminOwnershipRejectConfig struct {
	Repo     owner.Repository
	Factory  owner.Factory
	Events   owner.Events
	CqrsBase decorator.Base
}

func NewAdminOwnershipRejectHandler(config AdminOwnershipRejectConfig) AdminOwnershipRejectHandler {
	return decorator.ApplyCommandDecorators[AdminOwnershipRejectCommand, *AdminOwnershipRejectResult](
		&adminOwnershipRejectHandler{
			repo:    config.Repo,
			factory: config.Factory,
			events:  config.Events,
		},
		config.CqrsBase,
	)
}

func (h *adminOwnershipRejectHandler) Handle(ctx context.Context, cmd AdminOwnershipRejectCommand) (*AdminOwnershipRejectResult, *i18np.Error) {
	err := h.repo.Reject(ctx, cmd.OwnerNickName, cmd.Reason)
	if err != nil {
		return nil, err
	}
	h.events.RejectedByAdmin(&owner.EventOwnerRejectedByAdmin{
		OwnerNickName: cmd.OwnerNickName,
		AdminUUID:     cmd.AdminUUID,
		Reason:        cmd.Reason,
	})
	return &AdminOwnershipRejectResult{}, nil
}
