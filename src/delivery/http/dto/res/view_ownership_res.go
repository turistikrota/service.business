package res

import (
	"time"

	"github.com/turistikrota/service.owner/src/app/query"
)

type ViewOwnershipResponse struct {
	NickName   string     `json:"nickName"`
	RealName   string     `json:"realName"`
	OwnerType  string     `json:"ownerType"`
	IsVerified bool       `json:"isVerified"`
	CreatedAt  *time.Time `json:"createdAt"`
}

func (r *response) ViewOwnership(res *query.ViewOwnershipResult) *ViewOwnershipResponse {
	return &ViewOwnershipResponse{
		NickName:   res.Ownership.NickName,
		RealName:   res.Ownership.RealName,
		OwnerType:  string(res.Ownership.OwnerType),
		IsVerified: res.Ownership.IsVerified,
		CreatedAt:  res.Ownership.CreatedAt,
	}
}
