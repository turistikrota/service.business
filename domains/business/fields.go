package business

type fieldsType struct {
	UUID            string
	NickName        string
	RealName        string
	BusinessType    string
	Individual      string
	Corporation     string
	Users           string
	IsEnabled       string
	IsVerified      string
	IsDeleted       string
	RejectReason    string
	Application     string
	PreferredLocale string
	ActivatedAt     string
	DisabledAt      string
	VerifiedAt      string
	CreatedAt       string
	UpdatedAt       string
}

type userFieldsType struct {
	UUID   string
	Name   string
	Roles  string
	JoinAt string
}

type individualFieldsType struct {
	FirstName      string
	LastName       string
	IdentityNumber string
	Province       string
	District       string
	Address        string
	SerialNumber   string
	DateOfBirth    string
}

type corporationFieldsType struct {
	TaxNumber string
	Type      string
	Province  string
	District  string
	Address   string
	TaxOffice string
	Title     string
}

var fields = fieldsType{
	UUID:            "_id",
	NickName:        "nick_name",
	RealName:        "real_name",
	BusinessType:    "business_type",
	Individual:      "individual",
	Corporation:     "corporation",
	Users:           "users",
	IsEnabled:       "is_enabled",
	IsVerified:      "is_verified",
	IsDeleted:       "is_deleted",
	VerifiedAt:      "verified_at",
	Application:     "application",
	PreferredLocale: "preferred_locale",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
	ActivatedAt:     "activated_at",
	DisabledAt:      "disabled_at",
	RejectReason:    "reject_reason",
}

var userFields = userFieldsType{
	UUID:   "uuid",
	Name:   "name",
	Roles:  "roles",
	JoinAt: "join_at",
}

var individualFields = individualFieldsType{
	FirstName:      "first_name",
	LastName:       "last_name",
	IdentityNumber: "identity_number",
	Province:       "province",
	District:       "district",
	Address:        "address",
	SerialNumber:   "serial_number",
	DateOfBirth:    "date_of_birth",
}

var corporationFields = corporationFieldsType{
	TaxNumber: "tax_number",
	Type:      "type",
	Province:  "province",
	District:  "district",
	Address:   "address",
	TaxOffice: "tax_office",
	Title:     "title",
}

func corporationField(field string) string {
	return fields.Corporation + "." + field
}

func individualField(field string) string {
	return fields.Individual + "." + field
}

func userField(field string) string {
	return fields.Users + "." + field
}

func userFieldInArray(field string) string {
	return fields.Users + ".$." + field
}

func userArrayFieldInArray(field string) string {
	return fields.Users + ".$[]." + field
}
