package query

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.owner/src/domain/owner"
)

type GetWithUserOwnershipQuery struct {
	NickName string
	UserName string
	UserUUID string
}

type GetWithUserOwnershipResult struct {
	Ownership *owner.EntityWithUser
}

type GetWithUserOwnershipHandler decorator.QueryHandler[GetWithUserOwnershipQuery, *GetWithUserOwnershipResult]

type getWithUserOwnershipHandler struct {
	repo    owner.Repository
	factory owner.Factory
}

type GetWithUserOwnershipHandlerConfig struct {
	Repo     owner.Repository
	Factory  owner.Factory
	CqrsBase decorator.Base
}

func NewGetWithUserOwnershipHandler(config GetWithUserOwnershipHandlerConfig) GetWithUserOwnershipHandler {
	return decorator.ApplyQueryDecorators[GetWithUserOwnershipQuery, *GetWithUserOwnershipResult](
		&getWithUserOwnershipHandler{
			repo:    config.Repo,
			factory: config.Factory,
		},
		config.CqrsBase,
	)
}

func (h *getWithUserOwnershipHandler) Handle(ctx context.Context, query GetWithUserOwnershipQuery) (*GetWithUserOwnershipResult, *i18np.Error) {
	ownership, err := h.repo.GetWithUser(ctx, query.NickName, owner.UserDetail{
		Name: query.UserName,
		UUID: query.UserUUID,
	})
	if err != nil {
		return nil, err
	}
	return &GetWithUserOwnershipResult{Ownership: ownership}, nil
}
