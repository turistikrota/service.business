package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.business/domains/business"
)

type AdminBusinessVerifyCmd struct {
	BusinessName string `params:"nickName" validate:"required"`
	AdminUUID    string `params:"-"`
}

type AdminBusinessVerifyRes struct{}

type AdminBusinessVerifyHandler cqrs.HandlerFunc[AdminBusinessVerifyCmd, *AdminBusinessVerifyRes]

func NewAdminBusinessVerifyHandler(repo business.Repository, factory business.Factory, events business.Events) AdminBusinessVerifyHandler {
	return func(ctx context.Context, cmd AdminBusinessVerifyCmd) (*AdminBusinessVerifyRes, *i18np.Error) {
		err := repo.Verify(ctx, cmd.BusinessName)
		if err != nil {
			return nil, err
		}
		events.VerifiedByAdmin(&business.EventBusinessVerifiedByAdmin{
			AdminUUID:    cmd.AdminUUID,
			BusinessName: cmd.BusinessName,
		})
		return &AdminBusinessVerifyRes{}, nil
	}
}
