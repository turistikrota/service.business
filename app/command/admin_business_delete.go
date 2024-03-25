package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.business/domains/business"
)

type AdminBusinessDeleteCmd struct {
	BusinessName string
	AdminUUID    string
}

type AdminBusinessDeleteRes struct{}

type AdminBusinessDeleteHandler cqrs.HandlerFunc[AdminBusinessDeleteCmd, *AdminBusinessDeleteRes]

func NewAdminBusinessDeleteHandler(repo business.Repository, factory business.Factory, events business.Events) AdminBusinessDeleteHandler {
	return func(ctx context.Context, cmd AdminBusinessDeleteCmd) (*AdminBusinessDeleteRes, *i18np.Error) {
		err := repo.Delete(ctx, cmd.BusinessName)
		if err != nil {
			return nil, err
		}
		events.DeletedByAdmin(&business.EventBusinessDeletedByAdmin{
			AdminUUID:    cmd.AdminUUID,
			BusinessName: cmd.BusinessName,
		})
		return &AdminBusinessDeleteRes{}, nil
	}
}
