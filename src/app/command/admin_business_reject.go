package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.business/src/domain/business"
)

type AdminBusinessRejectCommand struct {
	BusinessNickName string
	AdminUUID        string
	Reason           string
}

type AdminBusinessRejectResult struct{}

type AdminBusinessRejectHandler decorator.CommandHandler[AdminBusinessRejectCommand, *AdminBusinessRejectResult]

type adminBusinessRejectHandler struct {
	repo    business.Repository
	factory business.Factory
	events  business.Events
}

type AdminBusinessRejectConfig struct {
	Repo     business.Repository
	Factory  business.Factory
	Events   business.Events
	CqrsBase decorator.Base
}

func NewAdminBusinessRejectHandler(config AdminBusinessRejectConfig) AdminBusinessRejectHandler {
	return decorator.ApplyCommandDecorators[AdminBusinessRejectCommand, *AdminBusinessRejectResult](
		&adminBusinessRejectHandler{
			repo:    config.Repo,
			factory: config.Factory,
			events:  config.Events,
		},
		config.CqrsBase,
	)
}

func (h *adminBusinessRejectHandler) Handle(ctx context.Context, cmd AdminBusinessRejectCommand) (*AdminBusinessRejectResult, *i18np.Error) {
	res, _err := h.repo.GetByNickName(ctx, cmd.BusinessNickName)
	if _err != nil {
		return nil, _err
	}
	err := h.repo.Reject(ctx, cmd.BusinessNickName, cmd.Reason)
	if err != nil {
		return nil, err
	}
	h.events.RejectedByAdmin(&business.EventBusinessRejectedByAdmin{
		BusinessUUID: res.UUID,
		AdminUUID:    cmd.AdminUUID,
		Reason:       cmd.Reason,
	})
	return &AdminBusinessRejectResult{}, nil
}
