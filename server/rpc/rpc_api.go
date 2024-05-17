package rpc

import (
	"context"

	"github.com/turistikrota/service.business/app/query"
	protos "github.com/turistikrota/service.business/assets/protos/business"
	"github.com/turistikrota/service.business/domains/business"
)

func (s srv) ListAsClaim(ctx context.Context, req *protos.BusinessListAsClaimRequest) (*protos.BusinessListAsClaimResult, error) {
	businesses, err := s.app.Queries.ListAsClaim(ctx, query.ListAsClaimQuery{
		UserUUID: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	list := make([]*protos.Business, len(businesses.Businesses))
	for i, dto := range businesses.Businesses {
		var user business.User
		for _, u := range dto.Users {
			if u.UUID == req.UserId {
				user = u
				break
			}
		}
		list[i] = &protos.Business{
			Uuid:        dto.UUID,
			AccountName: user.Name,
			NickName:    dto.NickName,
			Roles:       user.Roles,
		}
	}
	return &protos.BusinessListAsClaimResult{
		Business: list,
	}, nil
}
