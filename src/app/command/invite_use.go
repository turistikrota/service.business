package command

import (
	"context"
	"time"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.business/src/domain/business"
	"github.com/turistikrota/service.business/src/domain/invite"
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
	businessRepo    business.Repository
	businessFactory business.Factory
	repo            invite.Repository
	factory         invite.Factory
	events          invite.Events
}

type InviteUseConfig struct {
	Repo            invite.Repository
	BusinessRepo    business.Repository
	BusinessFactory business.Factory
	Factory         invite.Factory
	Events          invite.Events
	CqrsBase        decorator.Base
}

func NewInviteUseHandler(config InviteUseConfig) InviteUseHandler {
	return decorator.ApplyCommandDecorators[InviteUseCommand, *InviteUseResult](
		&inviteUseHandler{
			repo:            config.Repo,
			factory:         config.Factory,
			businessRepo:    config.BusinessRepo,
			events:          config.Events,
			businessFactory: config.BusinessFactory,
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
	_err := h.businessRepo.AddUser(ctx, res.BusinessNickName, h.businessFactory.NewUser(cmd.UserUUID, cmd.UserName))
	if _err != nil {
		return nil, _err
	}
	error := h.repo.Use(ctx, cmd.InviteUUID)
	if error != nil {
		return nil, error
	}
	h.events.Use(invite.InviteUseEvent{
		InviteUUID:   cmd.InviteUUID,
		UserUUID:     cmd.UserUUID,
		UserName:     cmd.UserName,
		UserEmail:    cmd.UserEmail,
		BusinessUUID: res.BusinessUUID,
	})
	return &InviteUseResult{}, nil
}
