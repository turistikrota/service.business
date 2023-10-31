package query

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.owner/src/domain/invite"
)

type InviteGetByOwnerUUIDQuery struct {
	OwnerUUID string
}

type InviteGetByOwnerUUIDResult struct {
	Invites []*invite.Entity
}

type InviteGetByOwnerUUIDHandler decorator.QueryHandler[InviteGetByOwnerUUIDQuery, *InviteGetByOwnerUUIDResult]

type inviteGetByOwnerUUIDHandler struct {
	repo    invite.Repository
	factory invite.Factory
}

type InviteGetByOwnerUUIDHandlerConfig struct {
	Repo     invite.Repository
	Factory  invite.Factory
	CqrsBase decorator.Base
}

func NewInviteGetByOwnerUUIDHandler(config InviteGetByOwnerUUIDHandlerConfig) InviteGetByOwnerUUIDHandler {
	return decorator.ApplyQueryDecorators[InviteGetByOwnerUUIDQuery, *InviteGetByOwnerUUIDResult](
		&inviteGetByOwnerUUIDHandler{
			repo:    config.Repo,
			factory: config.Factory,
		},
		config.CqrsBase,
	)
}

func (h *inviteGetByOwnerUUIDHandler) Handle(ctx context.Context, query InviteGetByOwnerUUIDQuery) (*InviteGetByOwnerUUIDResult, *i18np.Error) {
	invites, err := h.repo.GetByOwnerUUID(ctx, query.OwnerUUID)
	if err != nil {
		return nil, h.factory.Errors.Failed(err.Error())
	}
	return &InviteGetByOwnerUUIDResult{Invites: invites}, nil
}
