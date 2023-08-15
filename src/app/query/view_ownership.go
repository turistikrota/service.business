package query

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.owner/src/domain/owner"
)

type ViewOwnershipQuery struct {
	NickName string
}

type ViewOwnershipResult struct {
	Ownership *owner.Entity
}

type ViewOwnershipHandler decorator.QueryHandler[ViewOwnershipQuery, *ViewOwnershipResult]

type viewOwnershipHandler struct {
	repo    owner.Repository
	factory owner.Factory
}

type ViewOwnershipHandlerConfig struct {
	Repo     owner.Repository
	Factory  owner.Factory
	CqrsBase decorator.Base
}

func NewViewOwnershipHandler(config ViewOwnershipHandlerConfig) ViewOwnershipHandler {
	return decorator.ApplyQueryDecorators[ViewOwnershipQuery, *ViewOwnershipResult](
		&viewOwnershipHandler{
			repo:    config.Repo,
			factory: config.Factory,
		},
		config.CqrsBase,
	)
}

func (h *viewOwnershipHandler) Handle(ctx context.Context, cmd ViewOwnershipQuery) (*ViewOwnershipResult, *i18np.Error) {
	ownership, err := h.repo.ProfileView(ctx, cmd.NickName)
	if err != nil {
		return nil, err
	}
	return &ViewOwnershipResult{Ownership: ownership}, nil
}
