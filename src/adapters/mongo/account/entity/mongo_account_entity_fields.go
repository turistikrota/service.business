package entity

type fields struct {
	UserUUID   string
	UserName   string
	UserCode   string
	FullName   string
	Avatar     string
	IsActive   string
	IsDeleted  string
	IsVerified string
	BirthDate  string
	CreatedAt  string
}

var Fields = fields{
	UserUUID:   "user_uuid",
	UserName:   "user_name",
	UserCode:   "user_code",
	FullName:   "full_name",
	Avatar:     "avatar",
	IsActive:   "is_active",
	IsDeleted:  "is_deleted",
	IsVerified: "is_verified",
	BirthDate:  "birth_date",
	CreatedAt:  "created_at",
}
