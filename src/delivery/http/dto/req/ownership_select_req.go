package req

import "github.com/turistikrota/service.owner/src/app/query"

type OwnershipSelectRequest struct {
	NickName string `param:"nickName" validate:"required"`
}

func (r OwnershipSelectRequest) ToGetQuery(userUUID string, userName string) query.GetWithUserOwnershipQuery {
	return query.GetWithUserOwnershipQuery{
		NickName: r.NickName,
		UserName: userName,
		UserUUID: userUUID,
	}
}
