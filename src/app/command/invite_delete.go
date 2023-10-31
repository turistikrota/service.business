package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.owner/src/domain/invite"
)

type InviteDeleteCommand struct {
	InviteUUID string
}

type InviteDeleteResult struct{}

type InviteDeleteHandler decorator.CommandHandler[InviteDeleteCommand, *InviteDeleteResult]

type inviteDeleteHandler struct {
	repo    invite.Repository
	factory invite.Factory
}

type InviteDeleteConfig struct {
	Repo     invite.Repository
	Factory  invite.Factory
	CqrsBase decorator.Base
}

func NewInviteDeleteHandler(config InviteDeleteConfig) InviteDeleteHandler {
	return decorator.ApplyCommandDecorators[InviteDeleteCommand, *InviteDeleteResult](
		&inviteDeleteHandler{
			repo:    config.Repo,
			factory: config.Factory,
		},
		config.CqrsBase,
	)
}

func (h *inviteDeleteHandler) Handle(ctx context.Context, cmd InviteDeleteCommand) (*InviteDeleteResult, *i18np.Error) {
	err := h.repo.Delete(ctx, cmd.InviteUUID)
	if err != nil {
		return nil, err
	}
	return &InviteDeleteResult{}, nil
}
