package command

import (
	"context"
	"time"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.owner/src/domain/account"
)

type AccountUpdateCommand struct {
	UserUUID    string
	CurrentName string
	NewName     string
	FullName    string
	Avatar      string
	BirthDate   *time.Time
}
type AccountUpdateResult struct{}

type AccountUpdateHandler decorator.CommandHandler[AccountUpdateCommand, *AccountUpdateResult]

type accountUpdateHandler struct {
	repo account.Repository
}

type AccountUpdateHandlerConfig struct {
	Repo     account.Repository
	CqrsBase decorator.Base
}

func NewAccountUpdateHandler(config AccountUpdateHandlerConfig) AccountUpdateHandler {
	return decorator.ApplyCommandDecorators[AccountUpdateCommand, *AccountUpdateResult](
		&accountUpdateHandler{
			repo: config.Repo,
		},
		config.CqrsBase,
	)
}

func (h *accountUpdateHandler) Handle(ctx context.Context, cmd AccountUpdateCommand) (*AccountUpdateResult, *i18np.Error) {
	_ = h.repo.Update(ctx, account.UserUnique{
		UserUUID: cmd.UserUUID,
		Name:     cmd.CurrentName,
	}, &account.Entity{
		UserName:  cmd.NewName,
		FullName:  cmd.FullName,
		BirthDate: cmd.BirthDate,
	})
	return &AccountUpdateResult{}, nil
}
