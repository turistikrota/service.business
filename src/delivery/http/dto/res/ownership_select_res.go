package res

import (
	"github.com/turistikrota/service.owner/src/app/query"
	"github.com/turistikrota/service.owner/src/domain/owner"
)

type OwnershipSelectNotSelectedResponse struct {
	MustSelect bool `json:"mustSelect"`
}

type SelectOwnershipResponse struct {
	User  *owner.User   `json:"user"`
	Owner *owner.Entity `json:"owner"`
}

func (r *response) SelectOwnership(res *query.GetWithUserOwnershipResult) *SelectOwnershipResponse {
	return &SelectOwnershipResponse{
		User:  &res.Ownership.User,
		Owner: &res.Ownership.Entity,
	}
}

func (r *response) OwnershipSelectNotFound() *OwnershipSelectNotSelectedResponse {
	return &OwnershipSelectNotSelectedResponse{
		MustSelect: true,
	}
}
