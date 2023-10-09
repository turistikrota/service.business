package owner

import "time"

type Entity struct {
	UUID        string      `json:"uuid"`
	NickName    string      `json:"nickName"`
	RealName    string      `json:"realName"`
	AvatarURL   string      `json:"avatarURL"`
	CoverURL    string      `json:"coverURL"`
	OwnerType   Type        `json:"ownerType"`
	Individual  Individual  `json:"individual"`
	Corporation Corporation `json:"corporation"`
	Users       []User      `json:"users"`
	IsEnabled   bool        `json:"isEnabled"`
	IsVerified  bool        `json:"isVerified"`
	IsDeleted   bool        `json:"isDeleted"`
	VerifiedAt  *time.Time  `json:"verifiedAt"`
	CreatedAt   *time.Time  `json:"createdAt"`
	UpdatedAt   *time.Time  `json:"updatedAt"`
}

type AdminListDto struct {
	UUID       string `json:"uuid"`
	NickName   string `json:"nickName"`
	RealName   string `json:"realName"`
	OwnerType  string `json:"ownerType"`
	IsEnabled  bool   `json:"isEnabled"`
	IsVerified bool   `json:"isVerified"`
	IsDeleted  bool   `json:"isDeleted"`
	VerifiedAt string `json:"verifiedAt"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}

type EntityWithUser struct {
	Entity Entity
	User   User
}

type Type string

type ownerTypes struct {
	Individual  Type
	Corporation Type
}

var Types = ownerTypes{
	Individual:  "individual",
	Corporation: "corporation",
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
	TaxNumber []byte
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
	IdentityNumber []byte
	SerialNumber   []byte
	Province       string
	District       string
	Address        string
	DateOfBirth    time.Time
}

type User struct {
	UUID   string    `json:"uuid"`
	Name   string    `json:"name"`
	Code   string    `json:"code"`
	Roles  []string  `json:"roles"`
	JoinAt time.Time `json:"joinAt"`
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
