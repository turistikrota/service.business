package query

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.business/src/domain/business"
)

type AdminViewBusinessQuery struct {
	NickName string
}

type AdminViewBusinessResult struct {
	Business *business.Entity
}

type AdminViewBusinessHandler decorator.QueryHandler[AdminViewBusinessQuery, *AdminViewBusinessResult]

type adminViewBusinessHandler struct {
	repo    business.Repository
	factory business.Factory
}

type AdminViewBusinessHandlerConfig struct {
	Repo     business.Repository
	Factory  business.Factory
	CqrsBase decorator.Base
}

func NewAdminViewBusinessHandler(config AdminViewBusinessHandlerConfig) AdminViewBusinessHandler {
	return decorator.ApplyQueryDecorators[AdminViewBusinessQuery, *AdminViewBusinessResult](
		&adminViewBusinessHandler{
			repo:    config.Repo,
			factory: config.Factory,
		},
		config.CqrsBase,
	)
}

func (h *adminViewBusinessHandler) Handle(ctx context.Context, cmd AdminViewBusinessQuery) (*AdminViewBusinessResult, *i18np.Error) {
	business, err := h.repo.GetByNickName(ctx, cmd.NickName)
	if err != nil {
		return nil, err
	}
	return &AdminViewBusinessResult{Business: business}, nil
}
