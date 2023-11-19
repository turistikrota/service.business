package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.business/src/domain/business"
)

type AdminBusinessRecoverCommand struct {
	BusinessNickName string
	AdminUUID        string
}

type AdminBusinessRecoverResult struct{}

type AdminBusinessRecoverHandler decorator.CommandHandler[AdminBusinessRecoverCommand, *AdminBusinessRecoverResult]

type adminBusinessRecoverHandler struct {
	repo    business.Repository
	factory business.Factory
	events  business.Events
}

type AdminBusinessRecoverConfig struct {
	Repo     business.Repository
	Factory  business.Factory
	Events   business.Events
	CqrsBase decorator.Base
}

func NewAdminBusinessRecoverHandler(config AdminBusinessRecoverConfig) AdminBusinessRecoverHandler {
	return decorator.ApplyCommandDecorators[AdminBusinessRecoverCommand, *AdminBusinessRecoverResult](
		&adminBusinessRecoverHandler{
			repo:    config.Repo,
			factory: config.Factory,
			events:  config.Events,
		},
		config.CqrsBase,
	)
}

func (h *adminBusinessRecoverHandler) Handle(ctx context.Context, cmd AdminBusinessRecoverCommand) (*AdminBusinessRecoverResult, *i18np.Error) {
	res, _err := h.repo.GetByNickName(ctx, cmd.BusinessNickName)
	if _err != nil {
		return nil, _err
	}
	err := h.repo.Recover(ctx, cmd.BusinessNickName)
	if err != nil {
		return nil, err
	}
	h.events.RecoverByAdmin(&business.EventBusinessRecoverByAdmin{
		BusinessUUID: res.UUID,
		AdminUUID:    cmd.AdminUUID,
	})
	return &AdminBusinessRecoverResult{}, nil
}
