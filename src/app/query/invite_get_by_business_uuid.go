package query

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.business/src/domain/invite"
)

type InviteGetByBusinessUUIDQuery struct {
	BusinessUUID string
}

type InviteGetByBusinessUUIDResult struct {
	Invites []*invite.Entity
}

type InviteGetByBusinessUUIDHandler decorator.QueryHandler[InviteGetByBusinessUUIDQuery, *InviteGetByBusinessUUIDResult]

type inviteGetByBusinessUUIDHandler struct {
	repo    invite.Repository
	factory invite.Factory
}

type InviteGetByBusinessUUIDHandlerConfig struct {
	Repo     invite.Repository
	Factory  invite.Factory
	CqrsBase decorator.Base
}

func NewInviteGetByBusinessUUIDHandler(config InviteGetByBusinessUUIDHandlerConfig) InviteGetByBusinessUUIDHandler {
	return decorator.ApplyQueryDecorators[InviteGetByBusinessUUIDQuery, *InviteGetByBusinessUUIDResult](
		&inviteGetByBusinessUUIDHandler{
			repo:    config.Repo,
			factory: config.Factory,
		},
		config.CqrsBase,
	)
}

func (h *inviteGetByBusinessUUIDHandler) Handle(ctx context.Context, query InviteGetByBusinessUUIDQuery) (*InviteGetByBusinessUUIDResult, *i18np.Error) {
	invites, err := h.repo.GetByBusinessUUID(ctx, query.BusinessUUID)
	if err != nil {
		return nil, h.factory.Errors.Failed(err.Error())
	}
	return &InviteGetByBusinessUUIDResult{Invites: invites}, nil
}
