package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.business/domains/business"
)

type AdminBusinessRejectCmd struct {
	Reason       string `json:"reason" params:"-" validate:"required,min=1,max=500"`
	BusinessName string `params:"nickName" json:"-" validate:"required"`
	AdminUUID    string `params:"-" json:"-"`
}

type AdminBusinessRejectRes struct{}

type AdminBusinessRejectHandler cqrs.HandlerFunc[AdminBusinessRejectCmd, *AdminBusinessRejectRes]

func NewAdminBusinessRejectHandler(repo business.Repository, factory business.Factory, events business.Events) AdminBusinessRejectHandler {
	return func(ctx context.Context, cmd AdminBusinessRejectCmd) (*AdminBusinessRejectRes, *i18np.Error) {
		err := repo.Reject(ctx, cmd.BusinessName, cmd.Reason)
		if err != nil {
			return nil, err
		}
		events.RejectedByAdmin(&business.EventBusinessRejectedByAdmin{
			AdminUUID:    cmd.AdminUUID,
			BusinessName: cmd.BusinessName,
		})
		return &AdminBusinessRejectRes{}, nil
	}
}
