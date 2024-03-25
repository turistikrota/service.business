package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.business/domains/invite"
)

type InviteGetByUUIDQuery struct {
	UUID string `params:"uuid" validate:"required,object_id"`
}

type InviteGetByUUIDRes struct {
	Invite *invite.Entity
}

type InviteGetByUUIDHandler cqrs.HandlerFunc[InviteGetByUUIDQuery, *InviteGetByUUIDRes]

func NewInviteGetByUUIDHandler(repo invite.Repository, factory invite.Factory) InviteGetByUUIDHandler {
	return func(ctx context.Context, query InviteGetByUUIDQuery) (*InviteGetByUUIDRes, *i18np.Error) {
		invite, err := repo.GetByUUID(ctx, query.UUID)
		if err != nil {
			return nil, factory.Errors.Failed(err.Error())
		}
		return &InviteGetByUUIDRes{Invite: invite}, nil
	}
}
