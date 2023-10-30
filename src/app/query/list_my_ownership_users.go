package query

import (
	"context"
	"time"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.owner/protos/account"
	"github.com/turistikrota/service.owner/src/config"
	"github.com/turistikrota/service.owner/src/domain/owner"
	"github.com/turistikrota/service.shared/helper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type ListMyOwnershipUsersQuery struct {
	NickName string
	UserName string
}

type ListMyOwnershipUserType struct {
	Name       string    `json:"name"`
	FullName   string    `json:"fullName"`
	AvatarURL  string    `json:"avatarUrl"`
	Roles      []string  `json:"roles"`
	IsVerified bool      `json:"isVerified"`
	IsCurrent  bool      `json:"isCurrent"`
	JoinAt     time.Time `json:"joinAt"`
	BirthDate  time.Time `json:"birthDate"`
	CreatedAt  time.Time `json:"createdAt"`
}

type ListMyOwnershipUsersResult struct {
	Users []ListMyOwnershipUserType
}

type ListMyOwnershipUsersQueryHandler decorator.QueryHandler[ListMyOwnershipUsersQuery, *ListMyOwnershipUsersResult]

type listMyOwnershipUsersQueryHandler struct {
	ownerRepo    owner.Repository
	ownerFactory owner.Factory
	rpcConfig    config.Rpc
}

type ListMyOwnershipUsersQueryHandlerConfig struct {
	OwnerRepo    owner.Repository
	OwnerFactory owner.Factory
	CqrsBase     decorator.Base
	Rpc          config.Rpc
}

func NewListMyOwnershipUsersQueryHandler(config ListMyOwnershipUsersQueryHandlerConfig) ListMyOwnershipUsersQueryHandler {
	return decorator.ApplyQueryDecorators[ListMyOwnershipUsersQuery, *ListMyOwnershipUsersResult](
		&listMyOwnershipUsersQueryHandler{
			ownerRepo:    config.OwnerRepo,
			ownerFactory: config.OwnerFactory,
			rpcConfig:    config.Rpc,
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
	_users := make([]*account.UserUnique, 0, len(users))
	ownerUsers := make(map[string]owner.User)
	for _, user := range users {
		_users = append(_users, &account.UserUnique{
			Uuid: user.UUID,
			Name: user.Name,
		})
		ownerUsers[user.UUID] = user
	}
	accounts, error := h.callRpc(ctx, _users)
	if error != nil {
		return nil, h.ownerFactory.Errors.Failed("failed to get account" + error.Error())
	}
	for _, user := range accounts.Entities {
		ownerUser := ownerUsers[user.Uuid]
		e := ListMyOwnershipUserType{
			Name:       user.UserName,
			FullName:   user.FullName,
			AvatarURL:  helper.CDN.DressAvatar(user.UserName),
			Roles:      ownerUser.Roles,
			IsVerified: user.IsVerified,
			IsCurrent:  user.UserName == query.UserName,
			JoinAt:     ownerUser.JoinAt,
			CreatedAt:  user.CreatedAt.AsTime(),
		}
		if user.BirthDate != nil {
			e.BirthDate = user.BirthDate.AsTime()
		}
		res.Users = append(res.Users, e)
	}
	return res, nil
}

func (h *listMyOwnershipUsersQueryHandler) callRpc(ctx context.Context, users []*account.UserUnique) (*account.AccountListByIdsResult, error) {
	var opt grpc.DialOption
	if !h.rpcConfig.AccountUsesSsl {
		opt = grpc.WithTransportCredentials(insecure.NewCredentials())
	} else {
		opt = grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, ""))
	}
	conn, err := grpc.Dial(h.rpcConfig.AccountHost, opt)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := account.NewAccountServiceClient(conn)
	response, err := c.GetAccountListByIds(ctx, &account.GetAccountListByIdsRequest{
		Users: users,
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}
