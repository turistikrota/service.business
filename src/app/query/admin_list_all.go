package query

import (
	"context"
	"strconv"
	"time"

	"github.com/mixarchitecture/cache"
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
	List *list.Result[*owner.Entity]
}

type AdminListOwnershipHandler decorator.QueryHandler[AdminListOwnershipQuery, *AdminListOwnershipResult]

type adminListOwnershipHandler struct {
	repo    owner.Repository
	factory owner.Factory
	cache   cache.Client[*list.Result[*owner.Entity]]
}

type AdminListOwnershipHandlerConfig struct {
	Repo     owner.Repository
	Factory  owner.Factory
	CqrsBase decorator.Base
	CacheSrv cache.Service
}

func NewAdminListOwnershipHandler(config AdminListOwnershipHandlerConfig) AdminListOwnershipHandler {
	return decorator.ApplyQueryDecorators[AdminListOwnershipQuery, *AdminListOwnershipResult](
		&adminListOwnershipHandler{
			repo:    config.Repo,
			factory: config.Factory,
			cache:   cache.New[*list.Result[*owner.Entity]](config.CacheSrv),
		},
		config.CqrsBase,
	)
}

func (h *adminListOwnershipHandler) Handle(ctx context.Context, query AdminListOwnershipQuery) (*AdminListOwnershipResult, *i18np.Error) {
	cacheHandler := func() (*list.Result[*owner.Entity], *i18np.Error) {
		return h.repo.AdminListAll(ctx, list.Config{
			Offset: query.Offset,
			Limit:  query.Limit,
		})
	}
	res, err := h.cache.Creator(h.createCacheEntity).Handler(cacheHandler).Timeout(1*time.Minute).Get(ctx, h.generateCacheKey(query))
	if err != nil {
		return nil, err
	}
	return &AdminListOwnershipResult{
		List: res,
	}, nil
}

func (h adminListOwnershipHandler) createCacheEntity() *list.Result[*owner.Entity] {
	return &list.Result[*owner.Entity]{}
}

func (h adminListOwnershipHandler) generateCacheKey(query AdminListOwnershipQuery) string {
	return "admin_list_owners" + strconv.FormatInt(query.Offset, 10) + "_" + strconv.FormatInt(query.Limit, 10)
}
