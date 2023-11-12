package command

import (
	"context"
	"time"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.owner/src/domain/invite"
	"github.com/turistikrota/service.owner/src/domain/owner"
)

type InviteUseCommand struct {
	InviteUUID string
	UserEmail  string
	UserUUID   string
	UserName   string
}

type InviteUseResult struct{}

type InviteUseHandler decorator.CommandHandler[InviteUseCommand, *InviteUseResult]

type inviteUseHandler struct {
	ownerRepo    owner.Repository
	ownerFactory owner.Factory
	repo         invite.Repository
	factory      invite.Factory
	events       invite.Events
}

type InviteUseConfig struct {
	Repo         invite.Repository
	OwnerRepo    owner.Repository
	OwnerFactory owner.Factory
	Factory      invite.Factory
	Events       invite.Events
	CqrsBase     decorator.Base
}

func NewInviteUseHandler(config InviteUseConfig) InviteUseHandler {
	return decorator.ApplyCommandDecorators[InviteUseCommand, *InviteUseResult](
		&inviteUseHandler{
			repo:         config.Repo,
			factory:      config.Factory,
			ownerRepo:    config.OwnerRepo,
			events:       config.Events,
			ownerFactory: config.OwnerFactory,
		},
		config.CqrsBase,
	)
}

func (h *inviteUseHandler) Handle(ctx context.Context, cmd InviteUseCommand) (*InviteUseResult, *i18np.Error) {
	res, err := h.repo.GetByUUID(ctx, cmd.InviteUUID)
	if err != nil {
		return nil, err
	}
	if res.Email != cmd.UserEmail {
		return nil, h.factory.Errors.EmailMismatch()
	}
	if res.IsUsed {
		return nil, h.factory.Errors.Used()
	}
	if res.IsDeleted {
		return nil, h.factory.Errors.Deleted()
	}
	if res.CreatedAt.Add(24 * time.Hour).Before(time.Now()) {
		return nil, h.factory.Errors.Timeout()
	}
	_err := h.ownerRepo.AddUser(ctx, res.OwnerNickName, h.ownerFactory.NewUser(cmd.UserUUID, cmd.UserName))
	if _err != nil {
		return nil, _err
	}
	error := h.repo.Use(ctx, cmd.InviteUUID)
	if error != nil {
		return nil, error
	}
	h.events.Use(invite.InviteUseEvent{
		InviteUUID: cmd.InviteUUID,
		UserUUID:   cmd.UserUUID,
		UserName:   cmd.UserName,
		UserEmail:  cmd.UserEmail,
		OwnerUUID:  res.OwnerUUID,
	})
	return &InviteUseResult{}, nil
}
