package res

import (
	"time"

	"github.com/turistikrota/service.owner/src/domain/owner"
)

type OwnershipAdminViewResponse struct {
	NickName   string      `json:"nickName"`
	RealName   string      `json:"realName"`
	AvatarURL  string      `json:"avatarURL"`
	CoverURL   string      `json:"coverURL"`
	OwnerType  string      `json:"ownerType"`
	Details    interface{} `json:"details"`
	IsVerified bool        `json:"isVerified"`
	UpdatedAt  time.Time   `json:"updatedAt"`
	VerifiedAt time.Time   `json:"verifiedAt"`
	CreatedAt  time.Time   `json:"createdAt"`
}

type IndividualAdminViewResponse struct {
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Province  string    `json:"province"`
	District  string    `json:"district"`
	Address   string    `json:"address"`
	BirthDate time.Time `json:"birthDate"`
}

type CorporationAdminViewResponse struct {
	Province string `json:"province"`
	District string `json:"district"`
	Address  string `json:"address"`
	Type     string `json:"type"`
}

func (r *response) OwnershipAdminView(ownership *owner.Entity) *OwnershipAdminViewResponse {
	e := &OwnershipAdminViewResponse{
		NickName:   ownership.NickName,
		RealName:   ownership.RealName,
		AvatarURL:  ownership.AvatarURL,
		CoverURL:   ownership.CoverURL,
		OwnerType:  string(ownership.OwnerType),
		IsVerified: ownership.IsVerified,
		UpdatedAt:  *ownership.UpdatedAt,
		CreatedAt:  *ownership.CreatedAt,
	}
	if ownership.VerifiedAt != nil {
		e.VerifiedAt = *ownership.VerifiedAt
	}
	if ownership.OwnerType == owner.Types.Individual {
		e.Details = &CorporationAdminViewResponse{
			Province: ownership.Corporation.Province,
			District: ownership.Corporation.District,
			Address:  ownership.Corporation.Address,
			Type:     string(ownership.Corporation.Type),
		}
	} else {
		e.Details = &IndividualAdminViewResponse{
			FirstName: ownership.Individual.FirstName,
			LastName:  ownership.Individual.LastName,
			Province:  ownership.Individual.Province,
			District:  ownership.Individual.District,
			Address:   ownership.Individual.Address,
			BirthDate: ownership.Individual.DateOfBirth,
		}
	}
	return e
}
