package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.business/src/domain/business"
)

type BusinessEnableCommand struct {
	BusinessNickName string
	UserName         string
	UserCode         string
	UserUUID         string
}

type BusinessEnableResult struct{}

type BusinessEnableHandler decorator.CommandHandler[BusinessEnableCommand, *BusinessEnableResult]

type businessEnableHandler struct {
	repo    business.Repository
	factory business.Factory
	events  business.Events
}

type BusinessEnableConfig struct {
	Repo     business.Repository
	Factory  business.Factory
	Events   business.Events
	CqrsBase decorator.Base
}

func NewBusinessEnableHandler(config BusinessEnableConfig) BusinessEnableHandler {
	return decorator.ApplyCommandDecorators[BusinessEnableCommand, *BusinessEnableResult](
		&businessEnableHandler{
			repo:    config.Repo,
			factory: config.Factory,
			events:  config.Events,
		},
		config.CqrsBase,
	)
}

func (h *businessEnableHandler) Handle(ctx context.Context, cmd BusinessEnableCommand) (*BusinessEnableResult, *i18np.Error) {
	err := h.repo.Enable(ctx, cmd.BusinessNickName)
	if err != nil {
		return nil, err
	}
	h.events.Enabled(&business.EventBusinessEnabled{
		BusinessNickName: cmd.BusinessNickName,
		UserName:         cmd.UserName,
		UserCode:         cmd.UserCode,
		UserUUID:         cmd.UserUUID,
	})
	return &BusinessEnableResult{}, nil
}
