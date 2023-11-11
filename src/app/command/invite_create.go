package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.owner/src/domain/invite"
)

type InviteCreateCommand struct {
	OwnerNickName   string
	OwnerUUID       string
	CreatorUserName string
	Email           string
	Locale          string
	UserUUID        string
}

type InviteCreateResult struct{}

type InviteCreateHandler decorator.CommandHandler[InviteCreateCommand, *InviteCreateResult]

type inviteCreateHandler struct {
	repo    invite.Repository
	factory invite.Factory
	events  invite.Events
}

type InviteCreateConfig struct {
	Repo     invite.Repository
	Factory  invite.Factory
	Events   invite.Events
	CqrsBase decorator.Base
}

func NewInviteCreateHandler(config InviteCreateConfig) InviteCreateHandler {
	return decorator.ApplyCommandDecorators[InviteCreateCommand, *InviteCreateResult](
		&inviteCreateHandler{
			repo:    config.Repo,
			factory: config.Factory,
			events:  config.Events,
		},
		config.CqrsBase,
	)
}

func (h *inviteCreateHandler) Handle(ctx context.Context, cmd InviteCreateCommand) (*InviteCreateResult, *i18np.Error) {
	res, err := h.repo.Create(ctx, h.factory.New(cmd.Email, cmd.OwnerUUID, cmd.OwnerNickName, cmd.CreatorUserName))
	if err != nil {
		return nil, err
	}
	h.events.Invite(invite.InviteEvent{
		Locale:     cmd.Locale,
		Email:      cmd.Email,
		InviteUUID: res.UUID,
		OwnerName:  cmd.OwnerNickName,
		OwnerUUID:  cmd.OwnerUUID,
		UserUUID:   cmd.UserUUID,
		UserName:   cmd.CreatorUserName,
	})
	return &InviteCreateResult{}, nil
}
