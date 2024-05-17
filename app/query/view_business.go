package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.business/domains/business"
)

type ViewBusinessQuery struct {
	NickName string `params:"nickName" validate:"required"`
}

type ViewBusinessRes struct {
	Business *business.Entity
}

type ViewBusinessHandler cqrs.HandlerFunc[ViewBusinessQuery, *ViewBusinessRes]

func NewViewBusinessHandler(repo business.Repository) ViewBusinessHandler {
	return func(ctx context.Context, query ViewBusinessQuery) (*ViewBusinessRes, *i18np.Error) {
		b, err := repo.ProfileView(ctx, query.NickName)
		if err != nil {
			return nil, err
		}
		return &ViewBusinessRes{Business: b}, nil
	}
}
