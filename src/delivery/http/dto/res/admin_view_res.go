package res

import (
	"time"

	"github.com/turistikrota/service.business/src/app/query"
	"github.com/turistikrota/service.business/src/domain/business"
	"github.com/turistikrota/service.shared/helper"
)

type AdminViewRes struct {
	UUID         string                `json:"uuid"`
	NickName     string                `json:"nickName"`
	CoverURL     string                `json:"coverURL"`
	AvatarURL    string                `json:"avatarURL"`
	RealName     string                `json:"realName"`
	BusinessType string                `json:"businessType"`
	Individual   *AdminViewIndividual  `json:"individual,omitempty"`
	Corporation  *AdminViewCorporation `json:"corporation,omitempty"`
	Users        []business.User       `json:"users"`
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

func (r *response) AdminView(res *query.AdminViewBusinessResult) *AdminViewRes {
	rs := &AdminViewRes{
		UUID:         res.Business.UUID,
		NickName:     res.Business.NickName,
		RealName:     res.Business.RealName,
		AvatarURL:    helper.CDN.DressBusinessAvatar(res.Business.NickName),
		CoverURL:     helper.CDN.DressBusinessCover(res.Business.NickName),
		BusinessType: string(res.Business.BusinessType),
		Users:        res.Business.Users,
		RejectReason: res.Business.RejectReason,
		IsEnabled:    res.Business.IsEnabled,
		IsVerified:   res.Business.IsVerified,
		IsDeleted:    res.Business.IsDeleted,
		VerifiedAt:   res.Business.VerifiedAt,
		CreatedAt:    res.Business.CreatedAt,
		UpdatedAt:    res.Business.UpdatedAt,
	}
	if res.Business.BusinessType == business.Types.Individual {
		rs.Individual = &AdminViewIndividual{
			FirstName:   res.Business.Individual.FirstName,
			LastName:    res.Business.Individual.LastName,
			Province:    res.Business.Individual.Province,
			District:    res.Business.Individual.District,
			Address:     res.Business.Individual.Address,
			DateOfBirth: res.Business.Individual.DateOfBirth,
		}
	} else {
		rs.Corporation = &AdminViewCorporation{
			Province:  res.Business.Corporation.Province,
			District:  res.Business.Corporation.District,
			Address:   res.Business.Corporation.Address,
			TaxOffice: res.Business.Corporation.TaxOffice,
			Title:     res.Business.Corporation.Title,
			Type:      string(res.Business.Corporation.Type),
		}
	}
	return rs
}
