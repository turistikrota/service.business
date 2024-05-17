package business

import "time"

type Entity struct {
	UUID            string       `json:"uuid" bson:"_id,omitempty"`
	NickName        string       `json:"nickName" bson:"nick_name"`
	RealName        string       `json:"realName" bson:"real_name"`
	BusinessType    Type         `json:"businessType" bson:"business_type"`
	Individual      *Individual  `json:"individual,omitempty" bson:"individual,omitempty"`
	Corporation     *Corporation `json:"corporation,omitempty" bson:"corporation,omitempty"`
	Users           []User       `json:"users" bson:"users"`
	PreferredLocale Locale       `json:"preferredLocale" bson:"preferred_locale"`
	RejectReason    *string      `json:"rejectReason,omitempty" bson:"reject_reason,omitempty"`
	IsEnabled       bool         `json:"isEnabled" bson:"is_enabled"`
	IsVerified      bool         `json:"isVerified" bson:"is_verified"`
	IsDeleted       bool         `json:"isDeleted" bson:"is_deleted"`
	Application     Application  `json:"application" bson:"application"`
	VerifiedAt      *time.Time   `json:"verifiedAt" bson:"verified_at"`
	CreatedAt       *time.Time   `json:"createdAt" bson:"created_at"`
	UpdatedAt       *time.Time   `json:"updatedAt" bson:"updated_at"`
}

type Corporation struct {
	TaxNumber string          `json:"taxNumber" bson:"tax_number"`
	Province  string          `json:"province" bson:"province"`
	District  string          `json:"district" bson:"district"`
	Address   string          `json:"address" bson:"address"`
	TaxOffice string          `json:"taxOffice" bson:"tax_office"`
	Title     string          `json:"title" bson:"title"`
	Type      CorporationType `json:"type" bson:"type"`
}

type Individual struct {
	FirstName      string    `json:"firstName" bson:"first_name"`
	LastName       string    `json:"lastName" bson:"last_name"`
	IdentityNumber string    `json:"identityNumber" bson:"identity_number"`
	SerialNumber   string    `json:"serialNumber" bson:"serial_number"`
	Province       string    `json:"province" bson:"province"`
	District       string    `json:"district" bson:"district"`
	Address        string    `json:"address" bson:"address"`
	DateOfBirth    time.Time `json:"dateOfBirth" bson:"date_of_birth"`
}

type User struct {
	UUID   string    `json:"uuid" bson:"uuid"`
	Name   string    `json:"name" bson:"name"`
	Roles  []string  `json:"roles" bson:"roles"`
	JoinAt time.Time `json:"joinAt" bson:"join_at"`
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
