package res

import (
	"github.com/turistikrota/service.business/src/app/query"
)

func (r *response) ListMyBusinessUsers(res *query.ListMyBusinessUsersResult) []query.ListMyBusinessUserType {
	return res.Users
}
