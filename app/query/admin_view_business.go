package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.business/domains/business"
)

type AdminViewBusinessQuery struct {
	NickName string `params:"nickName" validate:"required"`
}

type AdminViewBusinessRes struct {
	Business *business.Entity
}

type AdminViewBusinessHandler cqrs.HandlerFunc[AdminViewBusinessQuery, *AdminViewBusinessRes]

func NewAdminViewBusinessHandler(repo business.Repository) AdminViewBusinessHandler {
	return func(ctx context.Context, query AdminViewBusinessQuery) (*AdminViewBusinessRes, *i18np.Error) {
		business, err := repo.GetByNickName(ctx, query.NickName)
		if err != nil {
			return nil, err
		}
		return &AdminViewBusinessRes{Business: business}, nil
	}
}
