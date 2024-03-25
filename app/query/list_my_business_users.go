package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.business/assets/protos/account"
	"github.com/turistikrota/service.business/config"
	"github.com/turistikrota/service.business/domains/business"
	"github.com/turistikrota/service.business/domains/user"
	"github.com/turistikrota/service.shared/helper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type ListMyBusinessUsersQuery struct {
	NickName string
	UserName string
}

type ListMyBusinessUsersRes struct {
	Users []user.ListByBusinessDto
}

type ListMyBusinessUsersHandler cqrs.HandlerFunc[ListMyBusinessUsersQuery, *ListMyBusinessUsersRes]

func NewListMyBusinessUsersHandler(repo business.Repository, factory business.Factory, rpcConfig config.Rpc) ListMyBusinessUsersHandler {

	callRpc := func(ctx context.Context, users []*account.UserUnique) (*account.AccountListByIdsResult, error) {
		var opt grpc.DialOption
		if !rpcConfig.AccountUsesSsl {
			opt = grpc.WithTransportCredentials(insecure.NewCredentials())
		} else {
			opt = grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, ""))
		}
		conn, err := grpc.Dial(rpcConfig.AccountHost, opt)
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

	return func(ctx context.Context, query ListMyBusinessUsersQuery) (*ListMyBusinessUsersRes, *i18np.Error) {
		users, err := repo.ListBusinessUsers(ctx, query.NickName, business.UserDetail{
			Name: query.UserName,
		})
		if err != nil {
			return nil, factory.Errors.Failed("failed to list business users")
		}
		res := &ListMyBusinessUsersRes{
			Users: make([]user.ListByBusinessDto, 0, len(users)),
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
		accounts, error := callRpc(ctx, _users)
		if error != nil {
			return nil, factory.Errors.Failed("failed to get account" + error.Error())
		}
		for _, u := range accounts.Entities {
			businessUser := businessUsers[u.UserName]
			e := user.ListByBusinessDto{
				Name:       u.UserName,
				FullName:   u.FullName,
				AvatarURL:  helper.CDN.DressAvatar(u.UserName),
				Roles:      businessUser.Roles,
				IsVerified: u.IsVerified,
				IsCurrent:  u.UserName == query.UserName,
				JoinAt:     businessUser.JoinAt,
				CreatedAt:  u.CreatedAt.AsTime(),
			}
			if u.BirthDate != nil {
				e.BirthDate = u.BirthDate.AsTime()
			}
			res.Users = append(res.Users, e)
		}
		return res, nil
	}
}
