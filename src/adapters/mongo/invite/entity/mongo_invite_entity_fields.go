package entity

type fields struct {
	UUID            string
	OwnerUUID       string
	OwnerName       string
	Email           string
	IsUsed          string
	IsDeleted       string
	CreatedAt       string
	UpdatedAt       string
	CreatorUserName string
}

var Fields = fields{
	UUID:            "_id",
	OwnerUUID:       "owner_uuid",
	OwnerName:       "owner_name",
	Email:           "email",
	IsUsed:          "is_used",
	CreatorUserName: "creator_user_name",
	IsDeleted:       "is_deleted",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
}
