package query

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.business/src/domain/business"
)

type ListMyBusinessesQuery struct {
	UserName string
	UserUUID string
}

type ListMyBusinessesResult struct {
	Businesses []*business.Entity
}

type ListMyBusinessesHandler decorator.QueryHandler[ListMyBusinessesQuery, *ListMyBusinessesResult]

type listMyBusinessesHandler struct {
	repo    business.Repository
	factory business.Factory
}

type ListMyBusinessesHandlerConfig struct {
	Repo     business.Repository
	Factory  business.Factory
	CqrsBase decorator.Base
}

func NewListMyBusinessesHandler(config ListMyBusinessesHandlerConfig) ListMyBusinessesHandler {
	return decorator.ApplyQueryDecorators[ListMyBusinessesQuery, *ListMyBusinessesResult](
		&listMyBusinessesHandler{
			repo:    config.Repo,
			factory: config.Factory,
		},
		config.CqrsBase,
	)
}

func (h *listMyBusinessesHandler) Handle(ctx context.Context, cmd ListMyBusinessesQuery) (*ListMyBusinessesResult, *i18np.Error) {
	businesss, err := h.repo.ListByUserUUID(ctx, business.UserDetail{
		Name: cmd.UserName,
		UUID: cmd.UserUUID,
	})
	if err != nil {
		return nil, h.factory.Errors.Failed(err.Error())
	}
	return &ListMyBusinessesResult{Businesses: businesss}, nil
}
