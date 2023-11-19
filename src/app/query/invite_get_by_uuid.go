package query

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.business/src/domain/invite"
)

type InviteGetByUUIDQuery struct {
	UUID string
}

type InviteGetByUUIDResult struct {
	Invite *invite.Entity
}

type InviteGetByUUIDHandler decorator.QueryHandler[InviteGetByUUIDQuery, *InviteGetByUUIDResult]

type inviteGetByUUIDHandler struct {
	repo    invite.Repository
	factory invite.Factory
}

type InviteGetByUUIDHandlerConfig struct {
	Repo     invite.Repository
	Factory  invite.Factory
	CqrsBase decorator.Base
}

func NewInviteGetByUUIDHandler(config InviteGetByUUIDHandlerConfig) InviteGetByUUIDHandler {
	return decorator.ApplyQueryDecorators[InviteGetByUUIDQuery, *InviteGetByUUIDResult](
		&inviteGetByUUIDHandler{
			repo:    config.Repo,
			factory: config.Factory,
		},
		config.CqrsBase,
	)
}

func (h *inviteGetByUUIDHandler) Handle(ctx context.Context, query InviteGetByUUIDQuery) (*InviteGetByUUIDResult, *i18np.Error) {
	invite, err := h.repo.GetByUUID(ctx, query.UUID)
	if err != nil {
		return nil, h.factory.Errors.Failed(err.Error())
	}
	return &InviteGetByUUIDResult{Invite: invite}, nil
}
