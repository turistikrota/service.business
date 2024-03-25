package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.business/domains/invite"
)

type InviteCreateCmd struct {
	BusinessName    string `json:"-"`
	BusinessUUID    string `json:"-"`
	CreatorUserName string `json:"-"`
	CreatorUserUUID string `json:"-"`
	Email           string `json:"email" validate:"required,email"`
	Locale          string `json:"locale" validate:"required,locale"`
}

type InviteCreateRes struct{}

type InviteCreateHandler cqrs.HandlerFunc[InviteCreateCmd, *InviteCreateRes]

func NewInviteCreateHandler(repo invite.Repository, factory invite.Factory, events invite.Events) InviteCreateHandler {
	return func(ctx context.Context, cmd InviteCreateCmd) (*InviteCreateRes, *i18np.Error) {
		res, err := repo.Create(ctx, factory.New(cmd.Email, cmd.BusinessUUID, cmd.BusinessName, cmd.CreatorUserName))
		if err != nil {
			return nil, err
		}
		events.Invite(invite.InviteEvent{
			Locale:       cmd.Locale,
			Email:        cmd.Email,
			InviteUUID:   res.UUID,
			BusinessName: cmd.BusinessName,
			BusinessUUID: cmd.BusinessUUID,
			UserUUID:     cmd.CreatorUserUUID,
			UserName:     cmd.CreatorUserName,
		})
		return &InviteCreateRes{}, nil
	}
}
