package res

import (
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
}

type response struct{}

func New() Response {
	return &response{}
}
