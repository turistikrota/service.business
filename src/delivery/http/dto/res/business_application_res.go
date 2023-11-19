package res

import "github.com/turistikrota/service.business/src/app/command"

type BusinessApplicationResponse struct {
	UUID string `json:"uuid"`
}

func (r *response) BusinessApplication(res *command.BusinessApplicationResult) *BusinessApplicationResponse {
	return &BusinessApplicationResponse{
		UUID: res.BusinessUUID,
	}
}
