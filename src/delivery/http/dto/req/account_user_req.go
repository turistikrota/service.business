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

func (r *AccountUserRequest) ToGetOwnershipQuery(nickName string, userUUID string) query.GetWithUserOwnershipQuery {
	return query.GetWithUserOwnershipQuery{
		NickName: nickName,
		UserName: r.CurrentUserName,
		UserUUID: userUUID,
	}
}
