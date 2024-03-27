package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.business/domains/business"
)

type ViewMyBusinessQuery struct {
	NickName string
}

type ViewMyBusinessRes struct {
	Business *business.Entity
}

type ViewMyBusinessHandler cqrs.HandlerFunc[ViewMyBusinessQuery, *ViewMyBusinessRes]

func NewViewMyBusinessHandler(repo business.Repository) ViewMyBusinessHandler {
	return func(ctx context.Context, query ViewMyBusinessQuery) (*ViewMyBusinessRes, *i18np.Error) {
		business, err := repo.GetByNickName(ctx, query.NickName)
		if err != nil {
			return nil, err
		}
		return &ViewMyBusinessRes{Business: business}, nil
	}
}
