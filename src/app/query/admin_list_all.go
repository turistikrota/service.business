package query

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/mixarchitecture/microp/types/list"
	"github.com/turistikrota/service.business/src/domain/business"
)

type AdminListBusinessQuery struct {
	Offset int64
	Limit  int64
}

type AdminListBusinessResult struct {
	List *list.Result[*business.AdminListDto]
}

type AdminListBusinessHandler decorator.QueryHandler[AdminListBusinessQuery, *AdminListBusinessResult]

type adminListBusinessHandler struct {
	repo    business.Repository
	factory business.Factory
}

type AdminListBusinessHandlerConfig struct {
	Repo     business.Repository
	Factory  business.Factory
	CqrsBase decorator.Base
}

func NewAdminListBusinessHandler(config AdminListBusinessHandlerConfig) AdminListBusinessHandler {
	return decorator.ApplyQueryDecorators[AdminListBusinessQuery, *AdminListBusinessResult](
		&adminListBusinessHandler{
			repo:    config.Repo,
			factory: config.Factory,
		},
		config.CqrsBase,
	)
}

func (h *adminListBusinessHandler) Handle(ctx context.Context, query AdminListBusinessQuery) (*AdminListBusinessResult, *i18np.Error) {
	res, err := h.repo.AdminListAll(ctx, list.Config{
		Offset: query.Offset,
		Limit:  query.Limit,
	})
	if err != nil {
		return nil, err
	}
	return &AdminListBusinessResult{
		List: res,
	}, nil
}
