package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.business/domains/invite"
)

type InviteGetByBusinessUUIDQuery struct {
	BusinessUUID string
}

type InviteGetByBusinessUUIDRes struct {
	Invites []*invite.Entity
}

type InviteGetByBusinessUUIDHandler cqrs.HandlerFunc[InviteGetByBusinessUUIDQuery, *InviteGetByBusinessUUIDRes]

func NewInviteGetByBusinessUUIDHandler(repo invite.Repository, factory invite.Factory) InviteGetByBusinessUUIDHandler {
	return func(ctx context.Context, query InviteGetByBusinessUUIDQuery) (*InviteGetByBusinessUUIDRes, *i18np.Error) {
		invites, err := repo.GetByBusinessUUID(ctx, query.BusinessUUID)
		if err != nil {
			return nil, factory.Errors.Failed(err.Error())
		}
		return &InviteGetByBusinessUUIDRes{Invites: invites}, nil
	}
}
