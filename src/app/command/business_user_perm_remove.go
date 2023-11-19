package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.business/src/domain/business"
)

type BusinessUserPermRemoveCommand struct {
	BusinessNickName string
	UserName         string
	AccessUserUUID   string
	AccessUserName   string
	PermissionName   string
}

type BusinessUserPermRemoveResult struct{}

type BusinessUserPermRemoveHandler decorator.CommandHandler[BusinessUserPermRemoveCommand, *BusinessUserPermRemoveResult]

type businessUserPermRemoveHandler struct {
	repo    business.Repository
	factory business.Factory
	events  business.Events
}

type BusinessUserPermRemoveHandlerConfig struct {
	Repo     business.Repository
	Factory  business.Factory
	Events   business.Events
	CqrsBase decorator.Base
}

func NewBusinessUserPermRemoveHandler(config BusinessUserPermRemoveHandlerConfig) BusinessUserPermRemoveHandler {
	return decorator.ApplyCommandDecorators[BusinessUserPermRemoveCommand, *BusinessUserPermRemoveResult](
		&businessUserPermRemoveHandler{
			repo:    config.Repo,
			factory: config.Factory,
			events:  config.Events,
		},
		config.CqrsBase,
	)
}

func (h *businessUserPermRemoveHandler) Handle(ctx context.Context, cmd BusinessUserPermRemoveCommand) (*BusinessUserPermRemoveResult, *i18np.Error) {
	res, _err := h.repo.GetWithUserName(ctx, cmd.BusinessNickName, cmd.UserName)
	if _err != nil {
		return nil, _err
	}
	err := h.repo.RemoveUserPermission(ctx, cmd.BusinessNickName, business.UserDetail{
		Name: cmd.UserName,
	}, cmd.PermissionName)
	if err != nil {
		return nil, err
	}
	h.events.UserPermissionRemoved(&business.EventBusinessPermissionRemoved{
		BusinessUUID:   res.Entity.UUID,
		AccessUserUUID: cmd.AccessUserUUID,
		AccessUserName: cmd.AccessUserName,
		PermissionName: cmd.PermissionName,
		User: business.EventUser{
			Name: cmd.UserName,
			UUID: res.User.UUID,
		},
	})
	return &BusinessUserPermRemoveResult{}, nil
}
