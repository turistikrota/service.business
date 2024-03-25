package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.business/domains/invite"
)

type InviteGetByEmailQuery struct {
	UserEmail string
}

type InviteGetByEmailRes struct {
	Invites []*invite.Entity
}

type InviteGetByEmailHandler cqrs.HandlerFunc[InviteGetByEmailQuery, *InviteGetByEmailRes]

func NewInviteGetByEmailHandler(repo invite.Repository, factory invite.Factory) InviteGetByEmailHandler {
	return func(ctx context.Context, query InviteGetByEmailQuery) (*InviteGetByEmailRes, *i18np.Error) {
		invites, err := repo.GetByEmail(ctx, query.UserEmail)
		if err != nil {
			return nil, factory.Errors.Failed(err.Error())
		}
		return &InviteGetByEmailRes{Invites: invites}, nil
	}
}
