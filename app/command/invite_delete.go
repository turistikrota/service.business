package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.business/domains/invite"
)

type InviteDeleteCmd struct {
	InviteUUID string `params:"uuid" validate:"required,uuid"`
	UserUUID   string `params:"-"`
	UserName   string `params:"-"`
}

type InviteDeleteRes struct{}

type InviteDeleteHandler cqrs.HandlerFunc[InviteDeleteCmd, *InviteDeleteRes]

func NewInviteDeleteHandler(repo invite.Repository, events invite.Events) InviteDeleteHandler {
	return func(ctx context.Context, cmd InviteDeleteCmd) (*InviteDeleteRes, *i18np.Error) {
		res, err := repo.GetByUUID(ctx, cmd.InviteUUID)
		if err != nil {
			return nil, err
		}
		_err := repo.Delete(ctx, cmd.InviteUUID)
		if _err != nil {
			return nil, _err
		}
		events.Delete(invite.InviteDeleteEvent{
			InviteUUID:       cmd.InviteUUID,
			UserUUID:         cmd.UserUUID,
			UserName:         cmd.UserName,
			BusinessNickName: res.BusinessName,
		})
		return &InviteDeleteRes{}, nil
	}
}
