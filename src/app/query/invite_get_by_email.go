package query

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.business/src/domain/invite"
)

type InviteGetByEmailQuery struct {
	UserEmail string
}

type InviteGetByEmailResult struct {
	Invites []*invite.Entity
}

type InviteGetByEmailHandler decorator.QueryHandler[InviteGetByEmailQuery, *InviteGetByEmailResult]

type inviteGetByEmailHandler struct {
	repo    invite.Repository
	factory invite.Factory
}

type InviteGetByEmailHandlerConfig struct {
	Repo     invite.Repository
	Factory  invite.Factory
	CqrsBase decorator.Base
}

func NewInviteGetByEmailHandler(config InviteGetByEmailHandlerConfig) InviteGetByEmailHandler {
	return decorator.ApplyQueryDecorators[InviteGetByEmailQuery, *InviteGetByEmailResult](
		&inviteGetByEmailHandler{
			repo:    config.Repo,
			factory: config.Factory,
		},
		config.CqrsBase,
	)
}

func (h *inviteGetByEmailHandler) Handle(ctx context.Context, query InviteGetByEmailQuery) (*InviteGetByEmailResult, *i18np.Error) {
	invites, err := h.repo.GetByEmail(ctx, query.UserEmail)
	if err != nil {
		return nil, h.factory.Errors.Failed(err.Error())
	}
	return &InviteGetByEmailResult{Invites: invites}, nil
}
