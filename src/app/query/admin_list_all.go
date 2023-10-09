package query

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/mixarchitecture/microp/types/list"
	"github.com/turistikrota/service.owner/src/domain/owner"
)

type AdminListOwnershipQuery struct {
	Offset int64
	Limit  int64
}

type AdminListOwnershipResult struct {
	List *list.Result[*owner.AdminListDto]
}

type AdminListOwnershipHandler decorator.QueryHandler[AdminListOwnershipQuery, *AdminListOwnershipResult]

type adminListOwnershipHandler struct {
	repo    owner.Repository
	factory owner.Factory
}

type AdminListOwnershipHandlerConfig struct {
	Repo     owner.Repository
	Factory  owner.Factory
	CqrsBase decorator.Base
}

func NewAdminListOwnershipHandler(config AdminListOwnershipHandlerConfig) AdminListOwnershipHandler {
	return decorator.ApplyQueryDecorators[AdminListOwnershipQuery, *AdminListOwnershipResult](
		&adminListOwnershipHandler{
			repo:    config.Repo,
			factory: config.Factory,
		},
		config.CqrsBase,
	)
}

func (h *adminListOwnershipHandler) Handle(ctx context.Context, query AdminListOwnershipQuery) (*AdminListOwnershipResult, *i18np.Error) {
	res, err := h.repo.AdminListAll(ctx, list.Config{
		Offset: query.Offset,
		Limit:  query.Limit,
	})
	if err != nil {
		return nil, err
	}
	return &AdminListOwnershipResult{
		List: res,
	}, nil
}
