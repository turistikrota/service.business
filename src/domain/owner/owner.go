package owner

import "time"

type Entity struct {
	UUID        string
	NickName    string
	RealName    string
	AvatarURL   string
	CoverURL    string
	OwnerType   Type
	Individual  Individual
	Corporation Corporation
	Users       []User
	IsEnabled   bool
	IsVerified  bool
	IsDeleted   bool
	VerifiedAt  *time.Time
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
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
	UUID   string
	Name   string
	Code   string
	Roles  []string
	JoinAt time.Time
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
