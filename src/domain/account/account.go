package account

import "time"

type Entity struct {
	UserUUID   string     `json:"user_uuid"`
	UserName   string     `json:"user_name"`
	UserCode   string     `json:"user_code"`
	FullName   string     `json:"full_name"`
	AvatarURL  string     `json:"avatar_url"`
	IsActive   bool       `json:"is_active"`
	IsDeleted  bool       `json:"is_deleted"`
	IsVerified bool       `json:"is_verified"`
	BirthDate  *time.Time `json:"birth_date"`
	CreatedAt  *time.Time `json:"created_at"`
}
