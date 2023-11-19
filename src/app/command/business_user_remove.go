package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.business/src/domain/business"
)

type BusinessUserRemoveCommand struct {
	BusinessNickName string
	UserName         string
	AccessUserUUID   string
	AccessUserName   string
}

type BusinessUserRemoveResult struct{}

type BusinessUserRemoveHandler decorator.CommandHandler[BusinessUserRemoveCommand, *BusinessUserRemoveResult]

type businessUserRemoveHandler struct {
	repo    business.Repository
	factory business.Factory
	events  business.Events
}

type BusinessUserRemoveHandlerConfig struct {
	Repo     business.Repository
	Factory  business.Factory
	Events   business.Events
	CqrsBase decorator.Base
}

func NewBusinessUserRemoveHandler(config BusinessUserRemoveHandlerConfig) BusinessUserRemoveHandler {
	return decorator.ApplyCommandDecorators[BusinessUserRemoveCommand, *BusinessUserRemoveResult](
		&businessUserRemoveHandler{
			repo:    config.Repo,
			factory: config.Factory,
			events:  config.Events,
		},
		config.CqrsBase,
	)
}

func (h *businessUserRemoveHandler) Handle(ctx context.Context, cmd BusinessUserRemoveCommand) (*BusinessUserRemoveResult, *i18np.Error) {
	res, _err := h.repo.GetWithUserName(ctx, cmd.BusinessNickName, cmd.UserName)
	if _err != nil {
		return nil, _err
	}
	err := h.repo.RemoveUser(ctx, cmd.BusinessNickName, business.UserDetail{
		Name: cmd.UserName,
	})
	if err != nil {
		return nil, h.factory.Errors.Failed("failed to remove user from business")
	}
	h.events.UserRemoved(&business.EventBusinessUserRemoved{
		BusinessUUID:   res.Entity.UUID,
		AccessUserUUID: cmd.AccessUserUUID,
		AccessUserName: cmd.AccessUserName,
		User: business.EventUser{
			Name: cmd.UserName,
			UUID: res.User.UUID,
		},
	})
	return &BusinessUserRemoveResult{}, nil
}
