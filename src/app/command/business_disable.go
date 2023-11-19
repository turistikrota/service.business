package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.business/src/domain/business"
)

type BusinessDisableCommand struct {
	BusinessNickName string
	UserName         string
	UserCode         string
	UserUUID         string
}

type BusinessDisableResult struct{}

type BusinessDisableHandler decorator.CommandHandler[BusinessDisableCommand, *BusinessDisableResult]

type businessDisableHandler struct {
	repo    business.Repository
	factory business.Factory
	events  business.Events
}

type BusinessDisableConfig struct {
	Repo     business.Repository
	Factory  business.Factory
	Events   business.Events
	CqrsBase decorator.Base
}

func NewBusinessDisableHandler(config BusinessDisableConfig) BusinessDisableHandler {
	return decorator.ApplyCommandDecorators[BusinessDisableCommand, *BusinessDisableResult](
		&businessDisableHandler{
			repo:    config.Repo,
			factory: config.Factory,
			events:  config.Events,
		},
		config.CqrsBase,
	)
}

func (h *businessDisableHandler) Handle(ctx context.Context, cmd BusinessDisableCommand) (*BusinessDisableResult, *i18np.Error) {
	res, _err := h.repo.GetByNickName(ctx, cmd.BusinessNickName)
	if _err != nil {
		return nil, _err
	}
	err := h.repo.Disable(ctx, cmd.BusinessNickName)
	if err != nil {
		return nil, err
	}

	h.events.Disabled(&business.EventBusinessDisabled{
		BusinessUUID: res.UUID,
		UserName:     cmd.UserName,
		UserCode:     cmd.UserCode,
		UserUUID:     cmd.UserUUID,
	})
	return &BusinessDisableResult{}, nil
}
