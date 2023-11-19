package res

import (
	"github.com/mixarchitecture/microp/types/list"
	"github.com/turistikrota/service.business/src/app/command"
	"github.com/turistikrota/service.business/src/app/query"
	"github.com/turistikrota/service.business/src/domain/business"
)

type Response interface {
	BusinessApplication(res *command.BusinessApplicationResult) *BusinessApplicationResponse
	BusinessAdminView(business *business.Entity) *BusinessAdminViewResponse
	ListMyBusinesses(res *query.ListMyBusinessesResult) *ListMyBusinessesResponse
	ListMyBusinessUsers(res *query.ListMyBusinessUsersResult) []query.ListMyBusinessUserType
	ViewBusiness(res *query.ViewBusinessResult) *ViewBusinessResponse
	SelectBusiness(res *query.GetWithUserBusinessResult) *SelectBusinessResponse
	BusinessSelectNotFound() *BusinessSelectNotSelectedResponse
	AdminListAll(res *query.AdminListBusinessResult) *list.Result[*business.AdminListDto]
	AdminView(res *query.AdminViewBusinessResult) *AdminViewRes
}

type response struct{}

func New() Response {
	return &response{}
}
