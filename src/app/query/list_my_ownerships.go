package query

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.owner/src/domain/owner"
)

type ListMyOwnershipsQuery struct {
	UserName string
	UserUUID string
}

type ListMyOwnershipsResult struct {
	Ownerships []*owner.Entity
}

type ListMyOwnershipsHandler decorator.QueryHandler[ListMyOwnershipsQuery, *ListMyOwnershipsResult]

type listMyOwnershipsHandler struct {
	repo    owner.Repository
	factory owner.Factory
}

type ListMyOwnershipsHandlerConfig struct {
	Repo     owner.Repository
	Factory  owner.Factory
	CqrsBase decorator.Base
}

func NewListMyOwnershipsHandler(config ListMyOwnershipsHandlerConfig) ListMyOwnershipsHandler {
	return decorator.ApplyQueryDecorators[ListMyOwnershipsQuery, *ListMyOwnershipsResult](
		&listMyOwnershipsHandler{
			repo:    config.Repo,
			factory: config.Factory,
		},
		config.CqrsBase,
	)
}

func (h *listMyOwnershipsHandler) Handle(ctx context.Context, cmd ListMyOwnershipsQuery) (*ListMyOwnershipsResult, *i18np.Error) {
	ownerships, err := h.repo.ListByUserUUID(ctx, owner.UserDetail{
		Name: cmd.UserName,
		UUID: cmd.UserUUID,
	})
	if err != nil {
		return nil, h.factory.Errors.Failed(err.Error())
	}
	return &ListMyOwnershipsResult{Ownerships: ownerships}, nil
}
