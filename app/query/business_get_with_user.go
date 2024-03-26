package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.business/domains/business"
)

type BusinessGetWithUserQuery struct {
	NickName string `params:"nickName" validate:"required"`
	UserName string `params:"-"`
	UserUUID string `params:"-"`
}

type BusinessGetWithUserRes struct {
	Dto *business.EntityWithUserDto
}

type BusinessGetWithUserHandler cqrs.HandlerFunc[BusinessGetWithUserQuery, *BusinessGetWithUserRes]

func NewBusinessGetWithUserHandler(repo business.Repository) BusinessGetWithUserHandler {
	return func(ctx context.Context, query BusinessGetWithUserQuery) (*BusinessGetWithUserRes, *i18np.Error) {
		business, err := repo.GetWithUser(ctx, query.NickName, business.UserDetail{
			Name: query.UserName,
			UUID: query.UserUUID,
		})
		if err != nil {
			return nil, err
		}
		return &BusinessGetWithUserRes{Dto: business}, nil
	}
}
