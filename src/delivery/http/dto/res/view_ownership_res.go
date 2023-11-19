package res

import (
	"time"

	"github.com/turistikrota/service.business/src/app/query"
)

type ViewBusinessResponse struct {
	NickName     string     `json:"nickName"`
	RealName     string     `json:"realName"`
	BusinessType string     `json:"businessType"`
	IsVerified   bool       `json:"isVerified"`
	CreatedAt    *time.Time `json:"createdAt"`
}

func (r *response) ViewBusiness(res *query.ViewBusinessResult) *ViewBusinessResponse {
	return &ViewBusinessResponse{
		NickName:     res.Business.NickName,
		RealName:     res.Business.RealName,
		BusinessType: string(res.Business.BusinessType),
		IsVerified:   res.Business.IsVerified,
		CreatedAt:    res.Business.CreatedAt,
	}
}
