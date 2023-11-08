package query

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.owner/src/domain/owner"
)

type AdminViewOwnershipQuery struct {
	NickName string
}

type AdminViewOwnershipResult struct {
	Ownership *owner.Entity
}

type AdminViewOwnershipHandler decorator.QueryHandler[AdminViewOwnershipQuery, *AdminViewOwnershipResult]

type adminViewOwnershipHandler struct {
	repo    owner.Repository
	factory owner.Factory
}

type AdminViewOwnershipHandlerConfig struct {
	Repo     owner.Repository
	Factory  owner.Factory
	CqrsBase decorator.Base
}

func NewAdminViewOwnershipHandler(config AdminViewOwnershipHandlerConfig) AdminViewOwnershipHandler {
	return decorator.ApplyQueryDecorators[AdminViewOwnershipQuery, *AdminViewOwnershipResult](
		&adminViewOwnershipHandler{
			repo:    config.Repo,
			factory: config.Factory,
		},
		config.CqrsBase,
	)
}

func (h *adminViewOwnershipHandler) Handle(ctx context.Context, cmd AdminViewOwnershipQuery) (*AdminViewOwnershipResult, *i18np.Error) {
	ownership, err := h.repo.GetByNickName(ctx, cmd.NickName)
	if err != nil {
		return nil, err
	}
	return &AdminViewOwnershipResult{Ownership: ownership}, nil
}
