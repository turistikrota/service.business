package entity

type fields struct {
	UUID        string
	NickName    string
	RealName    string
	AvatarURL   string
	CoverURL    string
	OwnerType   string
	Individual  string
	Corporation string
	Users       string
	IsEnabled   string
	IsVerified  string
	IsDeleted   string
	ActivatedAt string
	DisabledAt  string
	VerifiedAt  string
	CreatedAt   string
	UpdatedAt   string
}

type userFields struct {
	UUID   string
	Name   string
	Roles  string
	JoinAt string
}

type individualFields struct {
	FirstName      string
	LastName       string
	IdentityNumber string
	Province       string
	District       string
	Address        string
	SerialNumber   string
	DateOfBirth    string
}

type corporationFields struct {
	TaxNumber string
	Type      string
	Province  string
	District  string
	Address   string
	TaxOffice string
	Title     string
}

var Fields = fields{
	UUID:        "_id",
	NickName:    "nick_name",
	RealName:    "real_name",
	AvatarURL:   "avatar_url",
	CoverURL:    "cover_url",
	OwnerType:   "owner_type",
	Individual:  "individual",
	Corporation: "corporation",
	Users:       "users",
	IsEnabled:   "is_enabled",
	IsVerified:  "is_verified",
	IsDeleted:   "is_deleted",
	VerifiedAt:  "verified_at",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	ActivatedAt: "activated_at",
	DisabledAt:  "disabled_at",
}

var UserFields = userFields{
	UUID:   "uuid",
	Name:   "name",
	Roles:  "roles",
	JoinAt: "join_at",
}

var IndividualFields = individualFields{
	FirstName:      "first_name",
	LastName:       "last_name",
	IdentityNumber: "identity_number",
	Province:       "province",
	District:       "district",
	Address:        "address",
	SerialNumber:   "serial_number",
	DateOfBirth:    "date_of_birth",
}

var CorporationFields = corporationFields{
	TaxNumber: "tax_number",
	Type:      "type",
	Province:  "province",
	District:  "district",
	Address:   "address",
	TaxOffice: "tax_office",
	Title:     "title",
}

func CorporationField(field string) string {
	return Fields.Corporation + "." + field
}

func IndividualField(field string) string {
	return Fields.Individual + "." + field
}

func UserField(field string) string {
	return Fields.Users + "." + field
}

func UserFieldInArray(field string) string {
	return Fields.Users + ".$." + field
}

func UserArrayFieldInArray(field string) string {
	return Fields.Users + ".$[]." + field
}
