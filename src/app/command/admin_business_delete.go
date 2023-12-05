package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.business/src/domain/business"
)

type AdminBusinessDeleteCommand struct {
	BusinessNickName string
	AdminUUID        string
}

type AdminBusinessDeleteResult struct{}

type AdminBusinessDeleteHandler decorator.CommandHandler[AdminBusinessDeleteCommand, *AdminBusinessDeleteResult]

type adminBusinessDeleteHandler struct {
	repo    business.Repository
	factory business.Factory
	events  business.Events
}

type AdminBusinessDeleteConfig struct {
	Repo     business.Repository
	Factory  business.Factory
	Events   business.Events
	CqrsBase decorator.Base
}

func NewAdminBusinessDeleteHandler(config AdminBusinessDeleteConfig) AdminBusinessDeleteHandler {
	return decorator.ApplyCommandDecorators[AdminBusinessDeleteCommand, *AdminBusinessDeleteResult](
		&adminBusinessDeleteHandler{
			repo:    config.Repo,
			factory: config.Factory,
			events:  config.Events,
		},
		config.CqrsBase,
	)
}

func (h *adminBusinessDeleteHandler) Handle(ctx context.Context, cmd AdminBusinessDeleteCommand) (*AdminBusinessDeleteResult, *i18np.Error) {
	err := h.repo.Delete(ctx, cmd.BusinessNickName)
	if err != nil {
		return nil, err
	}
	h.events.DeletedByAdmin(&business.EventBusinessDeletedByAdmin{
		AdminUUID:        cmd.AdminUUID,
		BusinessNickName: cmd.BusinessNickName,
	})
	return &AdminBusinessDeleteResult{}, nil
}
