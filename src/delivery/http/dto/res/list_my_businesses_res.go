package res

import (
	"time"

	"github.com/turistikrota/service.business/src/app/query"
	"github.com/turistikrota/service.shared/helper"
)

type ListMyBusinessesResponse struct {
	List  []ListMyBusinessesItem `json:"list"`
	Count int                    `json:"count"`
}

type ListMyBusinessesItem struct {
	NickName     string    `json:"nickName"`
	RealName     string    `json:"realName"`
	CoverURL     string    `json:"coverURL"`
	AvatarURL    string    `json:"avatarURL"`
	BusinessType string    `json:"businessType"`
	RejectReason string    `json:"rejectReason,omitempty"`
	IsVerified   bool      `json:"isVerified"`
	IsEnabled    bool      `json:"isEnabled"`
	IsDeleted    bool      `json:"isDeleted"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

func (r *response) ListMyBusinesses(res *query.ListMyBusinessesResult) *ListMyBusinessesResponse {
	list := make([]ListMyBusinessesItem, 0)
	for _, business := range res.Businesses {
		item := ListMyBusinessesItem{
			NickName:     business.NickName,
			RealName:     business.RealName,
			AvatarURL:    helper.CDN.DressBusinessAvatar(business.NickName),
			CoverURL:     helper.CDN.DressBusinessCover(business.NickName),
			BusinessType: string(business.BusinessType),
			IsVerified:   business.IsVerified,
			IsEnabled:    business.IsEnabled,
			IsDeleted:    business.IsDeleted,
			UpdatedAt:    *business.UpdatedAt,
		}
		if business.RejectReason != nil {
			item.RejectReason = *business.RejectReason
		}
		list = append(list, item)
	}
	return &ListMyBusinessesResponse{
		List:  list,
		Count: len(list),
	}
}
