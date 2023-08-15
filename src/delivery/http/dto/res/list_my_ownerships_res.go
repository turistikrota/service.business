package res

import (
	"time"

	"github.com/turistikrota/service.owner/src/app/query"
)

type ListMyOwnershipsResponse struct {
	List  []ListMyOwnershipsItem `json:"list"`
	Count int                    `json:"count"`
}

type ListMyOwnershipsItem struct {
	NickName   string    `json:"nickName"`
	RealName   string    `json:"realName"`
	AvatarURL  string    `json:"avatarURL"`
	CoverURL   string    `json:"coverURL"`
	OwnerType  string    `json:"ownerType"`
	IsVerified bool      `json:"isVerified"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

func (r *response) ListMyOwnerships(res *query.ListMyOwnershipsResult) *ListMyOwnershipsResponse {
	list := make([]ListMyOwnershipsItem, 0)
	for _, ownership := range res.Ownerships {
		list = append(list, ListMyOwnershipsItem{
			NickName:   ownership.NickName,
			RealName:   ownership.RealName,
			AvatarURL:  ownership.AvatarURL,
			CoverURL:   ownership.CoverURL,
			OwnerType:  string(ownership.OwnerType),
			IsVerified: ownership.IsVerified,
			UpdatedAt:  *ownership.UpdatedAt,
		})
	}
	return &ListMyOwnershipsResponse{
		List:  list,
		Count: len(list),
	}
}
