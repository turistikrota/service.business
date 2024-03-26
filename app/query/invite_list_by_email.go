package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.business/domains/invite"
)

type InviteListByEmailQuery struct {
	UserEmail string
}

type InviteListByEmailRes struct {
	Invites []*invite.Entity
}

type InviteListByEmailHandler cqrs.HandlerFunc[InviteListByEmailQuery, *InviteListByEmailRes]

func NewInviteListByEmailHandler(repo invite.Repository, factory invite.Factory) InviteListByEmailHandler {
	return func(ctx context.Context, query InviteListByEmailQuery) (*InviteListByEmailRes, *i18np.Error) {
		invites, err := repo.ListByEmail(ctx, query.UserEmail)
		if err != nil {
			return nil, factory.Errors.Failed(err.Error())
		}
		return &InviteListByEmailRes{Invites: invites}, nil
	}
}
