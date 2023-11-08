package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.owner/src/domain/owner"
)

type AdminOwnershipDeleteCommand struct {
	OwnerNickName string
	AdminUUID     string
}

type AdminOwnershipDeleteResult struct{}

type AdminOwnershipDeleteHandler decorator.CommandHandler[AdminOwnershipDeleteCommand, *AdminOwnershipDeleteResult]

type adminOwnershipDeleteHandler struct {
	repo    owner.Repository
	factory owner.Factory
	events  owner.Events
}

type AdminOwnershipDeleteConfig struct {
	Repo     owner.Repository
	Factory  owner.Factory
	Events   owner.Events
	CqrsBase decorator.Base
}

func NewAdminOwnershipDeleteHandler(config AdminOwnershipDeleteConfig) AdminOwnershipDeleteHandler {
	return decorator.ApplyCommandDecorators[AdminOwnershipDeleteCommand, *AdminOwnershipDeleteResult](
		&adminOwnershipDeleteHandler{
			repo:    config.Repo,
			factory: config.Factory,
			events:  config.Events,
		},
		config.CqrsBase,
	)
}

func (h *adminOwnershipDeleteHandler) Handle(ctx context.Context, cmd AdminOwnershipDeleteCommand) (*AdminOwnershipDeleteResult, *i18np.Error) {
	err := h.repo.Delete(ctx, cmd.OwnerNickName)
	if err != nil {
		return nil, err
	}
	h.events.DeletedByAdmin(&owner.EventOwnerDeletedByAdmin{
		OwnerNickName: cmd.OwnerNickName,
		AdminUUID:     cmd.AdminUUID,
	})
	return &AdminOwnershipDeleteResult{}, nil
}
