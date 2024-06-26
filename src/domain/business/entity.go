package business

import "time"

type Entity struct {
	UUID            string      `json:"uuid"`
	NickName        string      `json:"nickName"`
	RealName        string      `json:"realName"`
	BusinessType    Type        `json:"businessType"`
	Individual      Individual  `json:"individual"`
	Corporation     Corporation `json:"corporation"`
	Users           []User      `json:"users"`
	PreferredLocale Locale      `json:"preferredLocale"`
	RejectReason    *string     `json:"rejectReason,omitempty"`
	IsEnabled       bool        `json:"isEnabled"`
	IsVerified      bool        `json:"isVerified"`
	IsDeleted       bool        `json:"isDeleted"`
	Application     Application `json:"application"`
	VerifiedAt      *time.Time  `json:"verifiedAt"`
	CreatedAt       *time.Time  `json:"createdAt"`
	UpdatedAt       *time.Time  `json:"updatedAt"`
}

type AdminListDto struct {
	UUID            string `json:"uuid"`
	NickName        string `json:"nickName"`
	RealName        string `json:"realName"`
	BusinessType    string `json:"businessType"`
	IsEnabled       bool   `json:"isEnabled"`
	IsVerified      bool   `json:"isVerified"`
	IsDeleted       bool   `json:"isDeleted"`
	Application     string `json:"application"`
	PreferredLocale Locale `json:"preferredLocale"`
	VerifiedAt      string `json:"verifiedAt,omitempty"`
	CreatedAt       string `json:"createdAt"`
	UpdatedAt       string `json:"updatedAt"`
}

type EntityWithUser struct {
	Entity Entity
	User   User
}

type Type string

type businessTypes struct {
	Individual  Type
	Corporation Type
}

var Types = businessTypes{
	Individual:  "individual",
	Corporation: "corporation",
}

type Application string

type applications struct {
	Accommodation Application
	Advert        Application
	Place         Application
}

var Applications = applications{
	Accommodation: "accommodation",
	Advert:        "advert",
	Place:         "place",
}

type CorporationType string

type corporationTypes struct {
	Anonymous                      CorporationType
	SoleProprietorship             CorporationType
	Limited                        CorporationType
	Collective                     CorporationType
	Cooperative                    CorporationType
	OrdinaryPartnership            CorporationType
	OrdinaryLimitedPartnership     CorporationType
	LimitedPartnershipShareCapital CorporationType
	Other                          CorporationType
}

var CorporationTypes = corporationTypes{
	Anonymous:                      "anonymous",
	SoleProprietorship:             "sole_proprietorship",
	Limited:                        "limited",
	Collective:                     "collective",
	Cooperative:                    "cooperative",
	OrdinaryPartnership:            "ordinary_partnership",
	OrdinaryLimitedPartnership:     "ordinary_limited_partnership",
	LimitedPartnershipShareCapital: "limited_partnership_share_capital",
	Other:                          "other",
}

type Corporation struct {
	TaxNumber string
	Province  string
	District  string
	Address   string
	TaxOffice string
	Title     string
	Type      CorporationType
}

type Individual struct {
	FirstName      string
	LastName       string
	IdentityNumber string
	SerialNumber   string
	Province       string
	District       string
	Address        string
	DateOfBirth    time.Time
}

type User struct {
	UUID   string    `json:"uuid"`
	Name   string    `json:"name"`
	Roles  []string  `json:"roles"`
	JoinAt time.Time `json:"joinAt"`
}

type Locale string

type locales struct {
	Tr Locale
	En Locale
}

var Locales = locales{
	Tr: "tr",
	En: "en",
}

func (u *User) HasPermission(permission string) bool {
	for _, role := range u.Roles {
		if role == permission {
			return true
		}
	}
	return false
}

func (u *EntityWithUser) HasPermissions(permissions ...string) bool {
	for _, permission := range permissions {
		if !u.User.HasPermission(permission) {
			return false
		}
	}
	return true
}

func (u *EntityWithUser) HasAnyPermissions(permissions ...string) bool {
	for _, permission := range permissions {
		if u.User.HasPermission(permission) {
			return true
		}
	}
	return false
}
