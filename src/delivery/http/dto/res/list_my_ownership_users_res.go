package res

import (
	"time"

	"github.com/turistikrota/service.owner/src/app/query"
)

type ListMyOwnershipUserItem struct {
	UserName   string     `json:"userName"`
	UserCode   string     `json:"userCode"`
	FullName   string     `json:"fullName"`
	AvatarURL  string     `json:"avatarURL"`
	Roles      []string   `json:"roles"`
	IsVerified bool       `json:"isVerified"`
	JoinAt     time.Time  `json:"joinAt"`
	BirthDate  *time.Time `json:"birthDate"`
	CreatedAt  *time.Time `json:"createdAt"`
}

func (r *response) ListMyOwnershipUsers(res *query.ListMyOwnershipUsersResult) []ListMyOwnershipUserItem {
	list := make([]ListMyOwnershipUserItem, 0)
	for _, user := range res.Users {
		list = append(list, ListMyOwnershipUserItem{
			UserName:   user.Name,
			UserCode:   user.Code,
			FullName:   user.FullName,
			AvatarURL:  user.AvatarURL,
			Roles:      user.Roles,
			IsVerified: user.IsVerified,
			JoinAt:     user.JoinAt,
			BirthDate:  user.BirthDate,
			CreatedAt:  user.CreatedAt,
		})
	}
	return list
}
