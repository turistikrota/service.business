package req

import "github.com/turistikrota/service.business/src/app/query"

type BusinessSelectRequest struct {
	NickName string `param:"nickName" validate:"required"`
}

func (r BusinessSelectRequest) ToGetQuery(userUUID string, userName string) query.GetWithUserBusinessQuery {
	return query.GetWithUserBusinessQuery{
		NickName: r.NickName,
		UserName: userName,
		UserUUID: userUUID,
	}
}
