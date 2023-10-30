package query

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.owner/protos/account"
	"github.com/turistikrota/service.owner/src/config"
	"github.com/turistikrota/service.owner/src/domain/owner"
	"github.com/turistikrota/service.shared/helper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type ListMyOwnershipUsersQuery struct {
	NickName string
	UserName string
}

type ListMyOwnershipUserType struct {
	Name       string               `json:"name"`
	FullName   string               `json:"fullName"`
	AvatarURL  string               `json:"avatarUrl"`
	Roles      []string             `json:"roles"`
	IsVerified bool                 `json:"isVerified"`
	JoinAt     time.Time            `json:"joinAt"`
	BirthDate  *timestamp.Timestamp `json:"birthDate"`
	CreatedAt  *timestamp.Timestamp `json:"createdAt"`
}

type ListMyOwnershipUsersResult struct {
	Users []ListMyOwnershipUserType
}

type ListMyOwnershipUsersQueryHandler decorator.QueryHandler[ListMyOwnershipUsersQuery, *ListMyOwnershipUsersResult]

type listMyOwnershipUsersQueryHandler struct {
	ownerRepo    owner.Repository
	ownerFactory owner.Factory
	hosts        config.RpcHosts
}

type ListMyOwnershipUsersQueryHandlerConfig struct {
	OwnerRepo    owner.Repository
	OwnerFactory owner.Factory
	CqrsBase     decorator.Base
	RpcHosts     config.RpcHosts
}

func NewListMyOwnershipUsersQueryHandler(config ListMyOwnershipUsersQueryHandlerConfig) ListMyOwnershipUsersQueryHandler {
	return decorator.ApplyQueryDecorators[ListMyOwnershipUsersQuery, *ListMyOwnershipUsersResult](
		&listMyOwnershipUsersQueryHandler{
			ownerRepo:    config.OwnerRepo,
			ownerFactory: config.OwnerFactory,
			hosts:        config.RpcHosts,
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
	ids := make([]string, 0, len(users))
	for _, user := range users {
		ids = append(ids, user.UUID)
	}
	accounts, error := h.callRpc(ctx, ids)
	if error != nil {
		return nil, h.ownerFactory.Errors.Failed("failed to get account" + error.Error())
	}
	for idx, user := range accounts.Entities {
		ownerUser := users[idx]
		res.Users = append(res.Users, ListMyOwnershipUserType{
			Name:       user.UserName,
			FullName:   user.FullName,
			AvatarURL:  helper.CDN.DressAvatar(user.UserName),
			Roles:      ownerUser.Roles,
			IsVerified: user.IsVerified,
			JoinAt:     ownerUser.JoinAt,
			BirthDate:  user.BirthDate,
			CreatedAt:  user.CreatedAt,
		})
	}
	return res, nil
}

func (h *listMyOwnershipUsersQueryHandler) callRpc(ctx context.Context, ids []string) (*account.AccountListByIdsResult, error) {
	conn, err := grpc.Dial(h.hosts.Account, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := account.NewAccountServiceClient(conn)
	response, err := c.GetAccountListByIds(ctx, &account.GetAccountListByIdsRequest{
		Uuids: ids,
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}
