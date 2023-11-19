package query

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.business/src/domain/business"
)

type GetWithUserBusinessQuery struct {
	NickName string
	UserName string
	UserUUID string
}

type GetWithUserBusinessResult struct {
	Business *business.EntityWithUser
}

type GetWithUserBusinessHandler decorator.QueryHandler[GetWithUserBusinessQuery, *GetWithUserBusinessResult]

type getWithUserBusinessHandler struct {
	repo    business.Repository
	factory business.Factory
}

type GetWithUserBusinessHandlerConfig struct {
	Repo     business.Repository
	Factory  business.Factory
	CqrsBase decorator.Base
}

func NewGetWithUserBusinessHandler(config GetWithUserBusinessHandlerConfig) GetWithUserBusinessHandler {
	return decorator.ApplyQueryDecorators[GetWithUserBusinessQuery, *GetWithUserBusinessResult](
		&getWithUserBusinessHandler{
			repo:    config.Repo,
			factory: config.Factory,
		},
		config.CqrsBase,
	)
}

func (h *getWithUserBusinessHandler) Handle(ctx context.Context, query GetWithUserBusinessQuery) (*GetWithUserBusinessResult, *i18np.Error) {
	business, err := h.repo.GetWithUser(ctx, query.NickName, business.UserDetail{
		Name: query.UserName,
		UUID: query.UserUUID,
	})
	if err != nil {
		return nil, err
	}
	return &GetWithUserBusinessResult{Business: business}, nil
}
