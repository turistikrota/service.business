package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.business/src/domain/business"
)

type AdminBusinessVerifyCommand struct {
	BusinessNickName string
	AdminUUID        string
}

type AdminBusinessVerifyResult struct{}

type AdminBusinessVerifyHandler decorator.CommandHandler[AdminBusinessVerifyCommand, *AdminBusinessVerifyResult]

type adminBusinessVerifyHandler struct {
	repo    business.Repository
	factory business.Factory
	events  business.Events
}

type AdminBusinessVerifyConfig struct {
	Repo     business.Repository
	Factory  business.Factory
	Events   business.Events
	CqrsBase decorator.Base
}

func NewAdminBusinessVerifyHandler(config AdminBusinessVerifyConfig) AdminBusinessVerifyHandler {
	return decorator.ApplyCommandDecorators[AdminBusinessVerifyCommand, *AdminBusinessVerifyResult](
		&adminBusinessVerifyHandler{
			repo:    config.Repo,
			factory: config.Factory,
			events:  config.Events,
		},
		config.CqrsBase,
	)
}

func (h *adminBusinessVerifyHandler) Handle(ctx context.Context, cmd AdminBusinessVerifyCommand) (*AdminBusinessVerifyResult, *i18np.Error) {
	err := h.repo.Verify(ctx, cmd.BusinessNickName)
	if err != nil {
		return nil, err
	}
	h.events.VerifiedByAdmin(&business.EventBusinessVerifiedByAdmin{
		BusinessNickName: cmd.BusinessNickName,
		AdminUUID:        cmd.AdminUUID,
	})
	return &AdminBusinessVerifyResult{}, nil
}
