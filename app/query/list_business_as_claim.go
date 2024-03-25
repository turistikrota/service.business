package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.business/domains/business"
)

type ListAsClaimQuery struct {
	UserUUID string
}

type ListAsClaimRes struct {
	Businesses []*business.Entity
}

type ListAsClaimHandler cqrs.HandlerFunc[ListAsClaimQuery, *ListAsClaimRes]

func NewListAsClaimHandler(repo business.Repository, factory business.Factory) ListAsClaimHandler {
	return func(ctx context.Context, query ListAsClaimQuery) (*ListAsClaimRes, *i18np.Error) {
		businesss, err := repo.ListAsClaim(ctx, query.UserUUID)
		if err != nil {
			return nil, factory.Errors.Failed(err.Error())
		}
		return &ListAsClaimRes{Businesses: businesss}, nil
	}
}
