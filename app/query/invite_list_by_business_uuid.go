package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.business/domains/invite"
)

type InviteListByBusinessUUIDQuery struct {
	BusinessUUID string
}

type InviteListByBusinessUUIDRes struct {
	Invites []*invite.Entity
}

type InviteListByBusinessUUIDHandler cqrs.HandlerFunc[InviteListByBusinessUUIDQuery, *InviteListByBusinessUUIDRes]

func NewInviteListByBusinessUUIDHandler(repo invite.Repository, factory invite.Factory) InviteListByBusinessUUIDHandler {
	return func(ctx context.Context, query InviteListByBusinessUUIDQuery) (*InviteListByBusinessUUIDRes, *i18np.Error) {
		invites, err := repo.ListByBusinessUUID(ctx, query.BusinessUUID)
		if err != nil {
			return nil, factory.Errors.Failed(err.Error())
		}
		return &InviteListByBusinessUUIDRes{Invites: invites}, nil
	}
}
