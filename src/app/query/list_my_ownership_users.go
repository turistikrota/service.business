package query

import (
	"context"
	"time"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.owner/src/domain/account"
	"github.com/turistikrota/service.owner/src/domain/owner"
	"github.com/turistikrota/service.shared/helper"
)

type ListMyOwnershipUsersQuery struct {
	NickName string
	UserName string
}

type ListMyOwnershipUserType struct {
	Name       string
	Code       string
	FullName   string
	AvatarURL  string
	Roles      []string
	IsVerified bool
	JoinAt     time.Time
	BirthDate  *time.Time
	CreatedAt  *time.Time
}

type ListMyOwnershipUsersResult struct {
	Users []ListMyOwnershipUserType
}

type ListMyOwnershipUsersQueryHandler decorator.QueryHandler[ListMyOwnershipUsersQuery, *ListMyOwnershipUsersResult]

type listMyOwnershipUsersQueryHandler struct {
	ownerRepo    owner.Repository
	accountRepo  account.Repository
	ownerFactory owner.Factory
}

type ListMyOwnershipUsersQueryHandlerConfig struct {
	OwnerRepo    owner.Repository
	AccountRepo  account.Repository
	OwnerFactory owner.Factory
	CqrsBase     decorator.Base
}

func NewListMyOwnershipUsersQueryHandler(config ListMyOwnershipUsersQueryHandlerConfig) ListMyOwnershipUsersQueryHandler {
	return decorator.ApplyQueryDecorators[ListMyOwnershipUsersQuery, *ListMyOwnershipUsersResult](
		&listMyOwnershipUsersQueryHandler{
			ownerRepo:    config.OwnerRepo,
			accountRepo:  config.AccountRepo,
			ownerFactory: config.OwnerFactory,
		},
		config.CqrsBase,
	)
}

func (h *listMyOwnershipUsersQueryHandler) Handle(ctx context.Context, query ListMyOwnershipUsersQuery) (*ListMyOwnershipUsersResult, *i18np.Error) {
	users, err := h.ownerRepo.ListOwnershipUsers(ctx, query.NickName, owner.UserDetail{
		Name: query.UserName,
	})
	if err != nil {
		return nil, h.ownerFactory.Errors.Failed("failed to list ownership users")
	}
	res := &ListMyOwnershipUsersResult{
		Users: make([]ListMyOwnershipUserType, 0, len(users)),
	}
	for _, user := range users {
		acc, err := h.accountRepo.Get(ctx, account.UserUnique{
			Name: user.Name,
		})
		if err != nil {
			return nil, h.ownerFactory.Errors.Failed("failed to get account")
		}
		res.Users = append(res.Users, ListMyOwnershipUserType{
			Name:       user.Name,
			Code:       user.Code,
			FullName:   acc.FullName,
			AvatarURL:  helper.CDN.DressAvatar(acc.UserName),
			Roles:      user.Roles,
			IsVerified: acc.IsVerified,
			JoinAt:     user.JoinAt,
			BirthDate:  acc.BirthDate,
			CreatedAt:  acc.CreatedAt,
		})
	}
	return res, nil
}
