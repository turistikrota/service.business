package query

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.business/src/domain/business"
)

type ViewBusinessQuery struct {
	NickName string
}

type ViewBusinessResult struct {
	Business *business.Entity
}

type ViewBusinessHandler decorator.QueryHandler[ViewBusinessQuery, *ViewBusinessResult]

type viewBusinessHandler struct {
	repo    business.Repository
	factory business.Factory
}

type ViewBusinessHandlerConfig struct {
	Repo     business.Repository
	Factory  business.Factory
	CqrsBase decorator.Base
}

func NewViewBusinessHandler(config ViewBusinessHandlerConfig) ViewBusinessHandler {
	return decorator.ApplyQueryDecorators[ViewBusinessQuery, *ViewBusinessResult](
		&viewBusinessHandler{
			repo:    config.Repo,
			factory: config.Factory,
		},
		config.CqrsBase,
	)
}

func (h *viewBusinessHandler) Handle(ctx context.Context, cmd ViewBusinessQuery) (*ViewBusinessResult, *i18np.Error) {
	business, err := h.repo.ProfileView(ctx, cmd.NickName)
	if err != nil {
		return nil, err
	}
	return &ViewBusinessResult{Business: business}, nil
}
