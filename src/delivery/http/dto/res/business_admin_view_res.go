package res

import (
	"time"

	"github.com/turistikrota/service.business/src/domain/business"
)

type BusinessAdminViewResponse struct {
	NickName     string      `json:"nickName"`
	RealName     string      `json:"realName"`
	BusinessType string      `json:"businessType"`
	Details      interface{} `json:"details"`
	IsVerified   bool        `json:"isVerified"`
	IsEnabled    bool        `json:"isEnabled"`
	IsDeleted    bool        `json:"isDeleted"`
	UpdatedAt    time.Time   `json:"updatedAt"`
	VerifiedAt   time.Time   `json:"verifiedAt"`
	CreatedAt    time.Time   `json:"createdAt"`
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

func (r *response) BusinessAdminView(b *business.Entity) *BusinessAdminViewResponse {
	e := &BusinessAdminViewResponse{
		NickName:     b.NickName,
		RealName:     b.RealName,
		BusinessType: string(b.BusinessType),
		IsVerified:   b.IsVerified,
		IsEnabled:    b.IsEnabled,
		IsDeleted:    b.IsDeleted,
		UpdatedAt:    *b.UpdatedAt,
		CreatedAt:    *b.CreatedAt,
	}
	if b.VerifiedAt != nil {
		e.VerifiedAt = *b.VerifiedAt
	}
	if b.BusinessType == business.Types.Individual {
		e.Details = &IndividualAdminViewResponse{
			FirstName: b.Individual.FirstName,
			LastName:  b.Individual.LastName,
			Province:  b.Individual.Province,
			District:  b.Individual.District,
			Address:   b.Individual.Address,
			BirthDate: b.Individual.DateOfBirth,
		}
	} else {
		e.Details = &CorporationAdminViewResponse{
			Province: b.Corporation.Province,
			District: b.Corporation.District,
			Address:  b.Corporation.Address,
			Type:     string(b.Corporation.Type),
		}
	}
	return e
}
