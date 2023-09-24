package owner

type messages struct {
	Failed                        string
	NickNameAlreadyExists         string
	IdentityNumberAlreadyExists   string
	TaxNumberAlreadyExists        string
	NotFound                      string
	TypeRequired                  string
	TypeInvalid                   string
	CorporationTypeRequired       string
	CorporationTypeInvalid        string
	IdentityVerificationFailed    string
	CorporationVerificationFailed string
	IndividualAlreadyExists       string
	CorporationAlreadyExists      string
}

var I18nMessages = messages{
	Failed:                        "error_owner_failed",
	NickNameAlreadyExists:         "error_owner_nickname_already_exists",
	IdentityNumberAlreadyExists:   "error_owner_identity_number_already_exists",
	TaxNumberAlreadyExists:        "error_owner_tax_number_already_exists",
	NotFound:                      "error_owner_not_found",
	TypeRequired:                  "error_owner_type_required",
	TypeInvalid:                   "error_owner_type_invalid",
	CorporationTypeRequired:       "error_owner_corporation_type_required",
	CorporationTypeInvalid:        "error_owner_corporation_type_invalid",
	IdentityVerificationFailed:    "error_owner_identity_verification_failed",
	CorporationVerificationFailed: "error_owner_corporation_verification_failed",
	IndividualAlreadyExists:       "error_owner_individual_already_exists",
	CorporationAlreadyExists:      "error_owner_corporation_already_exists",
}
