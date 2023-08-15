package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.owner/src/domain/owner"
)

type AdminOwnershipVerifyCommand struct {
	OwnerNickName string
	AdminUUID     string
	Reason        string
}

type AdminOwnershipVerifyResult struct{}

type AdminOwnershipVerifyHandler decorator.CommandHandler[AdminOwnershipVerifyCommand, *AdminOwnershipVerifyResult]

type adminOwnershipVerifyHandler struct {
	repo    owner.Repository
	factory owner.Factory
	events  owner.Events
}

type AdminOwnershipVerifyConfig struct {
	Repo     owner.Repository
	Factory  owner.Factory
	Events   owner.Events
	CqrsBase decorator.Base
}

func NewAdminOwnershipVerifyHandler(config AdminOwnershipVerifyConfig) AdminOwnershipVerifyHandler {
	return decorator.ApplyCommandDecorators[AdminOwnershipVerifyCommand, *AdminOwnershipVerifyResult](
		&adminOwnershipVerifyHandler{
			repo:    config.Repo,
			factory: config.Factory,
			events:  config.Events,
		},
		config.CqrsBase,
	)
}

func (h *adminOwnershipVerifyHandler) Handle(ctx context.Context, cmd AdminOwnershipVerifyCommand) (*AdminOwnershipVerifyResult, *i18np.Error) {
	err := h.repo.Verify(ctx, cmd.OwnerNickName)
	if err != nil {
		return nil, err
	}
	h.events.VerifiedByAdmin(&owner.EventOwnerVerifiedByAdmin{
		OwnerNickName: cmd.OwnerNickName,
		Reason:        cmd.Reason,
		AdminUUID:     cmd.AdminUUID,
	})
	return &AdminOwnershipVerifyResult{}, nil
}
