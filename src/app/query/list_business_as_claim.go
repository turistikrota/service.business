package query

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.business/src/domain/business"
)

type ListAsClaimQuery struct {
	UserUUID string
}

type ListAsClaimResult struct {
	Businesses []*business.Entity
}

type ListAsClaimHandler decorator.QueryHandler[ListAsClaimQuery, *ListAsClaimResult]

type listAsClaimHandler struct {
	repo    business.Repository
	factory business.Factory
}

type ListAsClaimHandlerConfig struct {
	Repo     business.Repository
	Factory  business.Factory
	CqrsBase decorator.Base
}

func NewListAsClaimHandler(config ListAsClaimHandlerConfig) ListAsClaimHandler {
	return decorator.ApplyQueryDecorators[ListAsClaimQuery, *ListAsClaimResult](
		&listAsClaimHandler{
			repo:    config.Repo,
			factory: config.Factory,
		},
		config.CqrsBase,
	)
}

func (h *listAsClaimHandler) Handle(ctx context.Context, cmd ListAsClaimQuery) (*ListAsClaimResult, *i18np.Error) {
	businesss, err := h.repo.ListAsClaim(ctx, cmd.UserUUID)
	if err != nil {
		return nil, h.factory.Errors.Failed(err.Error())
	}
	return &ListAsClaimResult{Businesses: businesss}, nil
}
