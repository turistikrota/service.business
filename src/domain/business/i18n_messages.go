package business

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
	ApplicationInvalid            string
	BusinessMustBeCorporation     string
	NotifySubjectRejected         string
	NotifySubjectVerified         string
	NotifyRejectContent           string
	NotifyVerifiedContent         string
}

var I18nMessages = messages{
	Failed:                        "error_business_failed",
	NickNameAlreadyExists:         "error_business_nickname_already_exists",
	IdentityNumberAlreadyExists:   "error_business_identity_number_already_exists",
	TaxNumberAlreadyExists:        "error_business_tax_number_already_exists",
	NotFound:                      "error_business_not_found",
	TypeRequired:                  "error_business_type_required",
	TypeInvalid:                   "error_business_type_invalid",
	CorporationTypeRequired:       "error_business_corporation_type_required",
	CorporationTypeInvalid:        "error_business_corporation_type_invalid",
	IdentityVerificationFailed:    "error_business_identity_verification_failed",
	CorporationVerificationFailed: "error_business_corporation_verification_failed",
	IndividualAlreadyExists:       "error_business_individual_already_exists",
	CorporationAlreadyExists:      "error_business_corporation_already_exists",
	ApplicationInvalid:            "error_business_application_invalid",
	BusinessMustBeCorporation:     "error_business_must_be_corporation",
	NotifySubjectRejected:         "notify_business_subject_rejected",
	NotifySubjectVerified:         "notify_business_subject_verified",
	NotifyRejectContent:           "notify_business_reject_content",
	NotifyVerifiedContent:         "notify_business_verify_content",
}
