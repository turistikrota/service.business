package req

import (
	"github.com/turistikrota/service.owner/src/app/query"
	"github.com/turistikrota/service.shared/helper"
)

type AccountUserRequest struct {
	UserName        string
	UserCode        string
	CurrentUserName string `params:"currentUserName" validate:"required,username_and_code"`
}

func (r *AccountUserRequest) Parse() *AccountUserRequest {
	_, r.UserName, r.UserCode = helper.Parsers.ParseUsernameAndCode(r.CurrentUserName)
	return r
}

func (r *AccountUserRequest) ToListMyOwnershipsQuery() query.ListMyOwnershipsQuery {
	r.Parse()
	return query.ListMyOwnershipsQuery{
		UserName: r.UserName,
		UserCode: r.UserCode,
	}
}
