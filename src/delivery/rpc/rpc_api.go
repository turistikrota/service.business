package rpc

import (
	"context"

	protos "github.com/turistikrota/service.business/protos/business"
	"github.com/turistikrota/service.business/src/app/query"
	"github.com/turistikrota/service.business/src/domain/business"
)

func (h Server) ListBusinessAsClaim(ctx context.Context, req *protos.BusinessListAsClaimRequest) (*protos.BusinessListAsClaimResult, error) {
	businesses, err := h.app.Queries.ListAsClaim.Handle(ctx, query.ListAsClaimQuery{
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
