package entity

type fields struct {
	UUID            string
	BusinessUUID    string
	BusinessName    string
	Email           string
	IsUsed          string
	IsDeleted       string
	CreatedAt       string
	UpdatedAt       string
	CreatorUserName string
}

var Fields = fields{
	UUID:            "_id",
	BusinessUUID:    "business_uuid",
	BusinessName:    "business_name",
	Email:           "email",
	IsUsed:          "is_used",
	CreatorUserName: "creator_user_name",
	IsDeleted:       "is_deleted",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
}
