package command

import (
	"context"
	"time"

	"github.com/turistikrota/service.owner/src/domain/account"

	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.owner/src/config"
	"github.com/turistikrota/service.owner/src/domain/owner"
	"github.com/turistikrota/service.shared/decorator"
)

type OwnershipUserAddCommand struct {
	OwnerNickName  string
	UserName       string
	UserCode       string
	AccessUserUUID string
}

type OwnershipUserAddResult struct{}

type OwnershipUserAddHandler decorator.CommandHandler[OwnershipUserAddCommand, *OwnershipUserAddResult]

type ownershipUserAddHandler struct {
	repo        owner.Repository
	accountRepo account.Repository
	factory     owner.Factory
	events      owner.Events
}

type OwnershipUserAddHandlerConfig struct {
	Repo        owner.Repository
	AccountRepo account.Repository
	Factory     owner.Factory
	Events      owner.Events
	CqrsBase    decorator.Base
}

func NewOwnershipUserAddHandler(config OwnershipUserAddHandlerConfig) OwnershipUserAddHandler {
	return decorator.ApplyCommandDecorators[OwnershipUserAddCommand, *OwnershipUserAddResult](
		&ownershipUserAddHandler{
			repo:        config.Repo,
			accountRepo: config.AccountRepo,
			factory:     config.Factory,
			events:      config.Events,
		},
		config.CqrsBase,
	)
}

func (h *ownershipUserAddHandler) Handle(ctx context.Context, cmd OwnershipUserAddCommand) (*OwnershipUserAddResult, *i18np.Error) {
	user, _err := h.accountRepo.Get(ctx, account.UserUnique{
		Name: cmd.UserName,
		Code: cmd.UserCode,
	})
	if _err != nil {
		return nil, _err
	}
	u := &owner.User{
		UUID:   user.UserUUID,
		Name:   cmd.UserName,
		Code:   cmd.UserCode,
		Roles:  []string{config.Roles.Owner.Member},
		JoinAt: time.Now(),
	}
	err := h.repo.AddUser(ctx, cmd.OwnerNickName, u)
	if err != nil {
		return nil, err
	}
	h.events.UserAdded(&owner.EventOwnerUserAdded{
		OwnerNickName: cmd.OwnerNickName,
		User:          u,
	})
	return &OwnershipUserAddResult{}, nil
}
