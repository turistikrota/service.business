package res

import (
	"github.com/mixarchitecture/microp/types/list"
	"github.com/turistikrota/service.owner/src/app/query"
	"github.com/turistikrota/service.owner/src/domain/owner"
)

func (r *response) AdminListAll(res *query.AdminListOwnershipResult) *list.Result[*owner.AdminListDto] {
	return res.List
}
