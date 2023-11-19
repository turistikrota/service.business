package res

import (
	"github.com/mixarchitecture/microp/types/list"
	"github.com/turistikrota/service.business/src/app/query"
	"github.com/turistikrota/service.business/src/domain/business"
)

func (r *response) AdminListAll(res *query.AdminListBusinessResult) *list.Result[*business.AdminListDto] {
	return res.List
}
