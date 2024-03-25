package business

type AdminListDto struct {
	UUID            string `json:"uuid"`
	NickName        string `json:"nickName"`
	RealName        string `json:"realName"`
	BusinessType    string `json:"businessType"`
	IsEnabled       bool   `json:"isEnabled"`
	IsVerified      bool   `json:"isVerified"`
	IsDeleted       bool   `json:"isDeleted"`
	Application     string `json:"application"`
	PreferredLocale Locale `json:"preferredLocale"`
	VerifiedAt      string `json:"verifiedAt,omitempty"`
	CreatedAt       string `json:"createdAt"`
	UpdatedAt       string `json:"updatedAt"`
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
