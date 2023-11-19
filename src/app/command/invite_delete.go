package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.business/src/domain/invite"
)

type InviteDeleteCommand struct {
	InviteUUID string
	UserUUID   string
	UserName   string
}

type InviteDeleteResult struct{}

type InviteDeleteHandler decorator.CommandHandler[InviteDeleteCommand, *InviteDeleteResult]

type inviteDeleteHandler struct {
	repo    invite.Repository
	factory invite.Factory
	events  invite.Events
}

type InviteDeleteConfig struct {
	Repo     invite.Repository
	Factory  invite.Factory
	CqrsBase decorator.Base
	Events   invite.Events
}

func NewInviteDeleteHandler(config InviteDeleteConfig) InviteDeleteHandler {
	return decorator.ApplyCommandDecorators[InviteDeleteCommand, *InviteDeleteResult](
		&inviteDeleteHandler{
			repo:    config.Repo,
			factory: config.Factory,
			events:  config.Events,
		},
		config.CqrsBase,
	)
}

func (h *inviteDeleteHandler) Handle(ctx context.Context, cmd InviteDeleteCommand) (*InviteDeleteResult, *i18np.Error) {
	res, err := h.repo.GetByUUID(ctx, cmd.InviteUUID)
	if err != nil {
		return nil, err
	}
	_err := h.repo.Delete(ctx, cmd.InviteUUID)
	if _err != nil {
		return nil, _err
	}
	h.events.Delete(invite.InviteDeleteEvent{
		InviteUUID:   cmd.InviteUUID,
		UserUUID:     cmd.UserUUID,
		UserName:     cmd.UserName,
		BusinessUUID: res.BusinessUUID,
	})
	return &InviteDeleteResult{}, nil
}
