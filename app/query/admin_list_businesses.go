package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/types/list"
	"github.com/turistikrota/service.business/domains/business"
	"github.com/turistikrota/service.business/pkg/paginate"
)

type AdminListBusinessesQuery struct {
	*paginate.Pagination
}

type AdminListBusinessesRes struct {
	List *list.Result[*business.AdminListDto]
}

type AdminListBusinessesHandler cqrs.HandlerFunc[AdminListBusinessesQuery, *AdminListBusinessesRes]

func NewAdminListBusinessesHandler(repo business.Repository) AdminListBusinessesHandler {
	return func(ctx context.Context, query AdminListBusinessesQuery) (*AdminListBusinessesRes, *i18np.Error) {
		query.Default()
		res, err := repo.AdminListAll(ctx, list.Config{
			Offset: (*query.Page - 1) * *query.Limit,
			Limit:  *query.Limit,
		})
		if err != nil {
			return nil, err
		}
		return &AdminListBusinessesRes{List: res}, nil
	}
}
