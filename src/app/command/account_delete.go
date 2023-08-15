package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.owner/src/domain/account"
)

type AccountDeleteCommand struct {
	UserUUID string
	Name     string
}
type AccountDeleteResult struct{}

type AccountDeleteHandler decorator.CommandHandler[AccountDeleteCommand, *AccountDeleteResult]

type accountDeleteHandler struct {
	repo account.Repository
}

type AccountDeleteHandlerConfig struct {
	Repo     account.Repository
	CqrsBase decorator.Base
}

func NewAccountDeleteHandler(config AccountDeleteHandlerConfig) AccountDeleteHandler {
	return decorator.ApplyCommandDecorators[AccountDeleteCommand, *AccountDeleteResult](
		&accountDeleteHandler{
			repo: config.Repo,
		},
		config.CqrsBase,
	)
}

func (h *accountDeleteHandler) Handle(ctx context.Context, cmd AccountDeleteCommand) (*AccountDeleteResult, *i18np.Error) {
	_ = h.repo.Delete(ctx, account.UserUnique{
		UserUUID: cmd.UserUUID,
		Name:     cmd.Name,
	})
	return &AccountDeleteResult{}, nil
}
