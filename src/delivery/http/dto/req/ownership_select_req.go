package req

import "github.com/turistikrota/service.owner/src/app/query"

type OwnershipSelectRequest struct {
	NickName string `param:"nickName" validate:"required"`
	UserName string `params:"currentUserName" validate:"required,username"`
}

func (r OwnershipSelectRequest) ToGetQuery(userUUID string) query.GetWithUserOwnershipQuery {
	return query.GetWithUserOwnershipQuery{
		NickName: r.NickName,
		UserName: r.UserName,
		UserUUID: userUUID,
	}
}
