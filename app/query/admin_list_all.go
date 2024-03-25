package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/types/list"
	"github.com/turistikrota/service.business/domains/business"
	"github.com/turistikrota/service.business/pkg/paginate"
)

type AdminListAllQuery struct {
	*paginate.Pagination
}

type AdminListAllRes struct {
	List *list.Result[*business.AdminListDto]
}

type AdminListAllHandler cqrs.HandlerFunc[AdminListAllQuery, *AdminListAllRes]

func NewAdminListAllHandler(repo business.Repository) AdminListAllHandler {
	return func(ctx context.Context, query AdminListAllQuery) (*AdminListAllRes, *i18np.Error) {
		query.Default()
		res, err := repo.AdminListAll(ctx, list.Config{
			Offset: (*query.Page - 1) * *query.Limit,
			Limit:  *query.Limit,
		})
		if err != nil {
			return nil, err
		}
		return &AdminListAllRes{List: res}, nil
	}
}
