package invite

import "time"

type Entity struct {
	UUID      string     `json:"uuid"`
	OwnerUUID string     `json:"ownerUUID"`
	Email     string     `json:"email"`
	IsUsed    bool       `json:"isUsed"`
	IsDeleted bool       `json:"isDeleted"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}
