package req

import (
	"github.com/turistikrota/service.owner/src/app/query"
)

type AccountUserRequest struct {
	CurrentUserName string `params:"currentUserName" validate:"required,username"`
}

func (r *AccountUserRequest) ToListMyOwnershipsQuery(userUUID string) query.ListMyOwnershipsQuery {
	return query.ListMyOwnershipsQuery{
		UserName: r.CurrentUserName,
		UserUUID: userUUID,
	}
}
