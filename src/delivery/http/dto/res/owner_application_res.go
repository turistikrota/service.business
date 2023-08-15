package res

import "github.com/turistikrota/service.owner/src/app/command"

type OwnerApplicationResponse struct {
	UUID string `json:"uuid"`
}

func (r *response) OwnerApplication(res *command.OwnerApplicationResult) *OwnerApplicationResponse {
	return &OwnerApplicationResponse{
		UUID: res.OwnerUUID,
	}
}
