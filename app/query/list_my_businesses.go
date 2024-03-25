package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.business/domains/business"
)

type ListMyBusinessesQuery struct {
	UserName string
	UserUUID string
}

type ListMyBusinessesRes struct {
	Businesses []*business.Entity
}

type ListMyBusinessesHandler cqrs.HandlerFunc[ListMyBusinessesQuery, *ListMyBusinessesRes]

func NewListMyBusinessesHandler(repo business.Repository, factory business.Factory) ListMyBusinessesHandler {
	return func(ctx context.Context, query ListMyBusinessesQuery) (*ListMyBusinessesRes, *i18np.Error) {
		businesss, err := repo.ListByUserUUID(ctx, business.UserDetail{
			Name: query.UserName,
			UUID: query.UserUUID,
		})
		if err != nil {
			return nil, factory.Errors.Failed(err.Error())
		}
		return &ListMyBusinessesRes{Businesses: businesss}, nil
	}
}
