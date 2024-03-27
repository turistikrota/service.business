package business

import "time"

type AdminListDto struct {
	UUID            string `json:"uuid"	bson:"uuid"`
	NickName        string `json:"nickName" bson:"nick_name"`
	RealName        string `json:"realName" bson:"real_name"`
	BusinessType    string `json:"businessType" bson:"business_type"`
	IsEnabled       bool   `json:"isEnabled" bson:"is_enabled"`
	IsVerified      bool   `json:"isVerified" bson:"is_verified"`
	IsDeleted       bool   `json:"isDeleted" bson:"is_deleted"`
	Application     string `json:"application" bson:"application"`
	PreferredLocale Locale `json:"preferredLocale" bson:"preferred_locale"`
	VerifiedAt      string `json:"verifiedAt,omitempty" bson:"verified_at"`
	CreatedAt       string `json:"createdAt" bson:"created_at"`
	UpdatedAt       string `json:"updatedAt" bson:"updated_at"`
}

type BusinessListDto struct {
	NickName     string     `json:"nickName" bson:"nick_name"`
	RealName     string     `json:"realName" bson:"real_name"`
	BusinessType string     `json:"businessType" bson:"business_type"`
	RejectReason *string    `json:"rejectReason,omitempty" bson:"reject_reason"`
	IsVerified   bool       `json:"isVerified" bson:"is_verified"`
	IsEnabled    bool       `json:"isEnabled" bson:"is_enabled"`
	IsDeleted    bool       `json:"isDeleted" bson:"is_deleted"`
	UpdatedAt    *time.Time `json:"updatedAt" bson:"updated_at"`
}

type EntityWithUserDto struct {
	Entity Entity
	User   User
}

func (u *EntityWithUserDto) HasPermissions(permissions ...string) bool {
	for _, permission := range permissions {
		if !u.User.HasPermission(permission) {
			return false
		}
	}
	return true
}

func (u *EntityWithUserDto) HasAnyPermissions(permissions ...string) bool {
	for _, permission := range permissions {
		if u.User.HasPermission(permission) {
			return true
		}
	}
	return false
}

func (e *Entity) ToEntityWithUserDto(u UserDetail) *EntityWithUserDto {
	for _, user := range e.Users {
		if user.Name == u.Name {
			return &EntityWithUserDto{
				Entity: *e,
				User:   user,
			}
		}
	}
	return nil
}
