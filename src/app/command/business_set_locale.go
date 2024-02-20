package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.business/src/domain/business"
)

type BusinessSetLocaleCommand struct {
	BusinessNickName string
	Locale           string
}

type BusinessSetLocaleResult struct{}

type BusinessSetLocaleHandler decorator.CommandHandler[BusinessSetLocaleCommand, *BusinessSetLocaleResult]

type businessSetLocaleHandler struct {
	repo    business.Repository
	factory business.Factory
	events  business.Events
}

type BusinessSetLocaleConfig struct {
	Repo     business.Repository
	Factory  business.Factory
	Events   business.Events
	CqrsBase decorator.Base
}

func NewBusinessSetLocaleHandler(config BusinessSetLocaleConfig) BusinessSetLocaleHandler {
	return decorator.ApplyCommandDecorators[BusinessSetLocaleCommand, *BusinessSetLocaleResult](
		&businessSetLocaleHandler{
			repo:    config.Repo,
			factory: config.Factory,
			events:  config.Events,
		},
		config.CqrsBase,
	)
}

func (h *businessSetLocaleHandler) Handle(ctx context.Context, cmd BusinessSetLocaleCommand) (*BusinessSetLocaleResult, *i18np.Error) {
	err := h.repo.SetPreferredLocale(ctx, cmd.BusinessNickName, cmd.Locale)
	if err != nil {
		return nil, err
	}
	return &BusinessSetLocaleResult{}, nil
}
