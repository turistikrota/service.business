package query

import (
	"context"
	"time"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.business/protos/account"
	"github.com/turistikrota/service.business/src/config"
	"github.com/turistikrota/service.business/src/domain/business"
	"github.com/turistikrota/service.shared/helper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type ListMyBusinessUsersQuery struct {
	NickName string
	UserName string
}

type ListMyBusinessUserType struct {
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

type ListMyBusinessUsersResult struct {
	Users []ListMyBusinessUserType
}

type ListMyBusinessUsersQueryHandler decorator.QueryHandler[ListMyBusinessUsersQuery, *ListMyBusinessUsersResult]

type listMyBusinessUsersQueryHandler struct {
	businessRepo    business.Repository
	businessFactory business.Factory
	rpcConfig       config.Rpc
}

type ListMyBusinessUsersQueryHandlerConfig struct {
	BusinessRepo    business.Repository
	BusinessFactory business.Factory
	CqrsBase        decorator.Base
	Rpc             config.Rpc
}

func NewListMyBusinessUsersQueryHandler(config ListMyBusinessUsersQueryHandlerConfig) ListMyBusinessUsersQueryHandler {
	return decorator.ApplyQueryDecorators[ListMyBusinessUsersQuery, *ListMyBusinessUsersResult](
		&listMyBusinessUsersQueryHandler{
			businessRepo:    config.BusinessRepo,
			businessFactory: config.BusinessFactory,
			rpcConfig:       config.Rpc,
		},
		config.CqrsBase,
	)
}

func (h *listMyBusinessUsersQueryHandler) Handle(ctx context.Context, query ListMyBusinessUsersQuery) (*ListMyBusinessUsersResult, *i18np.Error) {
	users, err := h.businessRepo.ListBusinessUsers(ctx, query.NickName, business.UserDetail{
		Name: query.UserName,
	})
	if err != nil {
		return nil, h.businessFactory.Errors.Failed("failed to list business users")
	}
	res := &ListMyBusinessUsersResult{
		Users: make([]ListMyBusinessUserType, 0, len(users)),
	}
	_users := make([]*account.UserUnique, 0, len(users))
	businessUsers := make(map[string]business.User)
	for _, user := range users {
		_users = append(_users, &account.UserUnique{
			Uuid: user.UUID,
			Name: user.Name,
		})
		businessUsers[user.Name] = user
	}
	accounts, error := h.callRpc(ctx, _users)
	if error != nil {
		return nil, h.businessFactory.Errors.Failed("failed to get account" + error.Error())
	}
	for _, user := range accounts.Entities {
		businessUser := businessUsers[user.UserName]
		e := ListMyBusinessUserType{
			Name:       user.UserName,
			FullName:   user.FullName,
			AvatarURL:  helper.CDN.DressAvatar(user.UserName),
			Roles:      businessUser.Roles,
			IsVerified: user.IsVerified,
			IsCurrent:  user.UserName == query.UserName,
			JoinAt:     businessUser.JoinAt,
			CreatedAt:  user.CreatedAt.AsTime(),
		}
		if user.BirthDate != nil {
			e.BirthDate = user.BirthDate.AsTime()
		}
		res.Users = append(res.Users, e)
	}
	return res, nil
}

func (h *listMyBusinessUsersQueryHandler) callRpc(ctx context.Context, users []*account.UserUnique) (*account.AccountListByIdsResult, error) {
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
