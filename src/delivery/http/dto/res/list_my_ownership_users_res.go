package res

import (
	"github.com/turistikrota/service.owner/src/app/query"
)

func (r *response) ListMyOwnershipUsers(res *query.ListMyOwnershipUsersResult) []query.ListMyOwnershipUserType {
	return res.Users
}
