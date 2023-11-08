package res

import (
	"github.com/mixarchitecture/microp/types/list"
	"github.com/turistikrota/service.owner/src/app/command"
	"github.com/turistikrota/service.owner/src/app/query"
	"github.com/turistikrota/service.owner/src/domain/owner"
)

type Response interface {
	OwnerApplication(res *command.OwnerApplicationResult) *OwnerApplicationResponse
	OwnershipAdminView(ownership *owner.Entity) *OwnershipAdminViewResponse
	ListMyOwnerships(res *query.ListMyOwnershipsResult) *ListMyOwnershipsResponse
	ListMyOwnershipUsers(res *query.ListMyOwnershipUsersResult) []query.ListMyOwnershipUserType
	ViewOwnership(res *query.ViewOwnershipResult) *ViewOwnershipResponse
	SelectOwnership(res *query.GetWithUserOwnershipResult) *SelectOwnershipResponse
	OwnershipSelectNotFound() *OwnershipSelectNotSelectedResponse
	AdminListAll(res *query.AdminListOwnershipResult) *list.Result[*owner.AdminListDto]
	AdminView(res *query.AdminViewOwnershipResult) *AdminViewRes
}

type response struct{}

func New() Response {
	return &response{}
}
