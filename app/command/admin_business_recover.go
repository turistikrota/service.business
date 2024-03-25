package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.business/domains/business"
)

type AdminBusinessRecoverCmd struct {
	BusinessName string
	AdminUUID    string
}

type AdminBusinessRecoverRes struct{}

type AdminBusinessRecoverHandler cqrs.HandlerFunc[AdminBusinessRecoverCmd, *AdminBusinessRecoverRes]

func NewAdminBusinessRecoverHandler(repo business.Repository, factory business.Factory, events business.Events) AdminBusinessRecoverHandler {
	return func(ctx context.Context, cmd AdminBusinessRecoverCmd) (*AdminBusinessRecoverRes, *i18np.Error) {
		err := repo.Recover(ctx, cmd.BusinessName)
		if err != nil {
			return nil, err
		}
		events.RecoverByAdmin(&business.EventBusinessRecoverByAdmin{
			AdminUUID:    cmd.AdminUUID,
			BusinessName: cmd.BusinessName,
		})
		return &AdminBusinessRecoverRes{}, nil
	}
}
