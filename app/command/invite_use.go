package command

import (
	"context"
	"time"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.business/domains/business"
	"github.com/turistikrota/service.business/domains/invite"
)

type InviteUseCmd struct {
	InviteUUID string `params:"uuid" validate:"required,uuid"`
	UserEmail  string `params:"-"`
	UserUUID   string `params:"-"`
	UserName   string `params:"-"`
}

type InviteUseRes struct{}

type InviteUseHandler cqrs.HandlerFunc[InviteUseCmd, *InviteUseRes]

func NewInviteUseHandler(repo invite.Repository, factory invite.Factory, events invite.Events, businessRepo business.Repository, businessFactory business.Factory) InviteUseHandler {
	return func(ctx context.Context, cmd InviteUseCmd) (*InviteUseRes, *i18np.Error) {
		res, err := repo.GetByUUID(ctx, cmd.InviteUUID)
		if err != nil {
			return nil, err
		}
		if res.Email != cmd.UserEmail {
			return nil, factory.Errors.EmailMismatch()
		}
		if res.IsUsed {
			return nil, factory.Errors.Used()
		}
		if res.IsDeleted {
			return nil, factory.Errors.Deleted()
		}
		if res.CreatedAt.Add(24 * time.Hour).Before(time.Now()) {
			return nil, factory.Errors.Timeout()
		}
		_err := businessRepo.AddUser(ctx, res.BusinessName, businessFactory.NewUser(cmd.UserUUID, cmd.UserName))
		if _err != nil {
			return nil, _err
		}
		error := repo.Use(ctx, cmd.InviteUUID)
		if error != nil {
			return nil, error
		}
		events.Use(invite.InviteUseEvent{
			InviteUUID:       cmd.InviteUUID,
			UserUUID:         cmd.UserUUID,
			UserName:         cmd.UserName,
			UserEmail:        cmd.UserEmail,
			BusinessNickName: res.BusinessName,
		})
		return &InviteUseRes{}, nil
	}
}
