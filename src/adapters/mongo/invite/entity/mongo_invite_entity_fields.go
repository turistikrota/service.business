package entity

type fields struct {
	UUID      string
	OwnerUUID string
	Email     string
	IsUsed    string
	IsDeleted string
	CreatedAt string
	UpdatedAt string
}

var Fields = fields{
	UUID:      "_id",
	OwnerUUID: "owner_uuid",
	Email:     "email",
	IsUsed:    "is_used",
	IsDeleted: "is_deleted",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}
