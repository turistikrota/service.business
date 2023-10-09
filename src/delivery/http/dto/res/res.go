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
	ListMyOwnershipUsers(res *query.ListMyOwnershipUsersResult) []ListMyOwnershipUserItem
	ViewOwnership(res *query.ViewOwnershipResult) *ViewOwnershipResponse
	SelectOwnership(res *query.GetWithUserOwnershipResult) *SelectOwnershipResponse
	OwnershipSelectNotFound() *OwnershipSelectNotSelectedResponse
	AdminListAll(res *query.AdminListOwnershipResult) *list.Result[*owner.AdminListDto]
}

type response struct{}

func New() Response {
	return &response{}
}
