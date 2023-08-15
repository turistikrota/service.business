package command

import (
	"context"
	"time"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.owner/src/domain/account"
)

type AccountCreateCommand struct {
	UserUUID    string
	AccountName string
	CreatedAt   *time.Time
}
type AccountCreateResult struct{}

type AccountCreateHandler decorator.CommandHandler[AccountCreateCommand, *AccountCreateResult]

type accountCreateHandler struct {
	repo account.Repository
}

type AccountCreateHandlerConfig struct {
	Repo     account.Repository
	CqrsBase decorator.Base
}

func NewAccountCreateHandler(config AccountCreateHandlerConfig) AccountCreateHandler {
	return decorator.ApplyCommandDecorators[AccountCreateCommand, *AccountCreateResult](
		&accountCreateHandler{
			repo: config.Repo,
		},
		config.CqrsBase,
	)
}

func (h *accountCreateHandler) Handle(ctx context.Context, cmd AccountCreateCommand) (*AccountCreateResult, *i18np.Error) {
	_ = h.repo.Create(ctx, &account.Entity{
		UserUUID:   cmd.UserUUID,
		UserName:   cmd.AccountName,
		FullName:   "",
		IsActive:   false,
		IsDeleted:  false,
		IsVerified: false,
		BirthDate:  nil,
		CreatedAt:  cmd.CreatedAt,
	})
	return &AccountCreateResult{}, nil
}
