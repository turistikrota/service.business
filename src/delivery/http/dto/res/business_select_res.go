package res

import (
	"github.com/turistikrota/service.business/src/app/query"
	"github.com/turistikrota/service.business/src/domain/business"
)

type BusinessSelectNotSelectedResponse struct {
	MustSelect bool `json:"mustSelect"`
}

type SelectBusinessResponse struct {
	User     *business.User   `json:"user"`
	Business *business.Entity `json:"business"`
}

func (r *response) SelectBusiness(res *query.GetWithUserBusinessResult) *SelectBusinessResponse {
	return &SelectBusinessResponse{
		User:     &res.Business.User,
		Business: &res.Business.Entity,
	}
}

func (r *response) BusinessSelectNotFound() *BusinessSelectNotSelectedResponse {
	return &BusinessSelectNotSelectedResponse{
		MustSelect: true,
	}
}
