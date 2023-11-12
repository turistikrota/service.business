package res

import (
	"time"

	"github.com/turistikrota/service.owner/src/app/query"
	"github.com/turistikrota/service.owner/src/domain/owner"
	"github.com/turistikrota/service.shared/helper"
)

type AdminViewRes struct {
	UUID         string                `json:"uuid"`
	NickName     string                `json:"nickName"`
	CoverURL     string                `json:"coverURL"`
	AvatarURL    string                `json:"avatarURL"`
	RealName     string                `json:"realName"`
	OwnerType    string                `json:"ownerType"`
	Individual   *AdminViewIndividual  `json:"individual,omitempty"`
	Corporation  *AdminViewCorporation `json:"corporation,omitempty"`
	Users        []owner.User          `json:"users"`
	RejectReason *string               `json:"rejectReason,omitempty"`
	IsEnabled    bool                  `json:"isEnabled"`
	IsVerified   bool                  `json:"isVerified"`
	IsDeleted    bool                  `json:"isDeleted"`
	VerifiedAt   *time.Time            `json:"verifiedAt"`
	CreatedAt    *time.Time            `json:"createdAt"`
	UpdatedAt    *time.Time            `json:"updatedAt"`
}

type AdminViewIndividual struct {
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Province    string    `json:"province"`
	District    string    `json:"district"`
	Address     string    `json:"address"`
	DateOfBirth time.Time `json:"dateOfBirth"`
}

type AdminViewCorporation struct {
	Province  string `json:"province"`
	District  string `json:"district"`
	Address   string `json:"address"`
	TaxOffice string `json:"taxOffice"`
	Title     string `json:"title"`
	Type      string `json:"type"`
}

func (r *response) AdminView(res *query.AdminViewOwnershipResult) *AdminViewRes {
	rs := &AdminViewRes{
		UUID:         res.Ownership.UUID,
		NickName:     res.Ownership.NickName,
		RealName:     res.Ownership.RealName,
		AvatarURL:    helper.CDN.DressOwnerAvatar(res.Ownership.NickName),
		CoverURL:     helper.CDN.DressOwnerCover(res.Ownership.NickName),
		OwnerType:    string(res.Ownership.OwnerType),
		Users:        res.Ownership.Users,
		RejectReason: res.Ownership.RejectReason,
		IsEnabled:    res.Ownership.IsEnabled,
		IsVerified:   res.Ownership.IsVerified,
		IsDeleted:    res.Ownership.IsDeleted,
		VerifiedAt:   res.Ownership.VerifiedAt,
		CreatedAt:    res.Ownership.CreatedAt,
		UpdatedAt:    res.Ownership.UpdatedAt,
	}
	if res.Ownership.OwnerType == owner.Types.Individual {
		rs.Individual = &AdminViewIndividual{
			FirstName:   res.Ownership.Individual.FirstName,
			LastName:    res.Ownership.Individual.LastName,
			Province:    res.Ownership.Individual.Province,
			District:    res.Ownership.Individual.District,
			Address:     res.Ownership.Individual.Address,
			DateOfBirth: res.Ownership.Individual.DateOfBirth,
		}
	} else {
		rs.Corporation = &AdminViewCorporation{
			Province:  res.Ownership.Corporation.Province,
			District:  res.Ownership.Corporation.District,
			Address:   res.Ownership.Corporation.Address,
			TaxOffice: res.Ownership.Corporation.TaxOffice,
			Title:     res.Ownership.Corporation.Title,
			Type:      string(res.Ownership.Corporation.Type),
		}
	}
	return rs
}
